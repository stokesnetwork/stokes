#!/bin/bash
# Script to build release binaries locally for Stokes

set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== Stokes Local Release Builder ===${NC}\n"

# Get version from user
if [ -z "$1" ]; then
    echo -e "${YELLOW}Usage: ./create-release.sh <version>${NC}"
    echo -e "Example: ./create-release.sh v0.1.0-testnet"
    echo -e "Example: ./create-release.sh v1.0.0"
    exit 1
fi

VERSION=$1

# Validate version format
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)?$ ]]; then
    echo -e "${RED}Error: Invalid version format${NC}"
    echo -e "Version should be in format: v0.1.0 or v0.1.0-testnet"
    exit 1
fi

echo -e "${YELLOW}Building release: $VERSION${NC}\n"

# Create releases directory
RELEASE_DIR="releases/$VERSION"
rm -rf "$RELEASE_DIR"
mkdir -p "$RELEASE_DIR"

# Platforms to build
declare -a PLATFORMS=(
    "linux:amd64"
    "linux:arm64"
    "darwin:amd64"
    "darwin:arm64"
    "windows:amd64"
)

echo -e "${BLUE}Building binaries for all platforms...${NC}\n"

for platform in "${PLATFORMS[@]}"; do
    IFS=':' read -r GOOS GOARCH <<< "$platform"
    
    echo -e "${YELLOW}Building $GOOS/$GOARCH...${NC}"
    
    # Create platform directory
    PLATFORM_NAME="$GOOS-$GOARCH"
    PLATFORM_DIR="$RELEASE_DIR/$PLATFORM_NAME"
    mkdir -p "$PLATFORM_DIR"
    
    # Determine binary suffix and C compiler
    SUFFIX=""
    CC_COMPILER=""
    case "$GOOS-$GOARCH" in
        linux-amd64)
            CC_COMPILER="x86_64-linux-musl-gcc"
            ;;
        linux-arm64)
            CC_COMPILER="aarch64-linux-musl-gcc"
            ;;
        windows-amd64)
            CC_COMPILER="x86_64-w64-mingw32-gcc"
            SUFFIX=".exe"
            ;;
        darwin-amd64|darwin-arm64)
            # Native macOS cross-compilation works without CC
            CC_COMPILER=""
            ;;
    esac
    
    # Build all binaries (CGO_ENABLED=1 for go-muhash compatibility)
    echo -n "  Building stokesd... "
    if CC=$CC_COMPILER CGO_ENABLED=1 GOOS=$GOOS GOARCH=$GOARCH go build -o "$PLATFORM_DIR/stokesd${SUFFIX}" . 2>&1; then
        echo -e "${GREEN}✓${NC}"
    else
        echo -e "${RED}✗ FAILED${NC}"
        exit 1
    fi
    
    echo -n "  Building stokesctl... "
    if CC=$CC_COMPILER CGO_ENABLED=1 GOOS=$GOOS GOARCH=$GOARCH go build -o "$PLATFORM_DIR/stokesctl${SUFFIX}" ./cmd/stokesctl 2>&1; then
        echo -e "${GREEN}✓${NC}"
    else
        echo -e "${RED}✗ FAILED${NC}"
        exit 1
    fi
    
    echo -n "  Building stokesminer... "
    if CC=$CC_COMPILER CGO_ENABLED=1 GOOS=$GOOS GOARCH=$GOARCH go build -o "$PLATFORM_DIR/stokesminer${SUFFIX}" ./cmd/stokesminer 2>&1; then
        echo -e "${GREEN}✓${NC}"
    else
        echo -e "${RED}✗ FAILED${NC}"
        exit 1
    fi
    
    echo -n "  Building stokeswallet... "
    if CC=$CC_COMPILER CGO_ENABLED=1 GOOS=$GOOS GOARCH=$GOARCH go build -o "$PLATFORM_DIR/stokeswallet${SUFFIX}" ./cmd/stokeswallet 2>&1; then
        echo -e "${GREEN}✓${NC}"
    else
        echo -e "${RED}✗ FAILED${NC}"
        exit 1
    fi
    
    # Generate platform-specific README and copy LICENSE
    ./generate-platform-readme.sh "$PLATFORM_NAME" "$VERSION" "$PLATFORM_DIR/README.md"
    cp LICENSE "$PLATFORM_DIR/"
    
    # Create archive
    echo -n "  Creating archive... "
    cd "$RELEASE_DIR"
    if [ "$GOOS" = "windows" ]; then
        zip -q -r "stokes-$VERSION-$PLATFORM_NAME.zip" "$PLATFORM_NAME"
        echo -e "${GREEN}✓${NC} stokes-$VERSION-$PLATFORM_NAME.zip"
    else
        tar -czf "stokes-$VERSION-$PLATFORM_NAME.tar.gz" "$PLATFORM_NAME"
        echo -e "${GREEN}✓${NC} stokes-$VERSION-$PLATFORM_NAME.tar.gz"
    fi
    cd - > /dev/null
    
    echo ""
done

echo -e "${GREEN}✓ All binaries built successfully!${NC}\n"
echo -e "${BLUE}Release files are in: ${YELLOW}$RELEASE_DIR${NC}\n"

# List all archives
echo -e "${BLUE}Archives created:${NC}"
ls -lh "$RELEASE_DIR"/*.{tar.gz,zip} 2>/dev/null | awk '{print "  " $9 " (" $5 ")"}'

echo -e "\n${GREEN}=== Next Steps: Upload to GitHub ===${NC}\n"
echo -e "${YELLOW}1. Create a new release on GitHub:${NC}"
echo -e "   https://github.com/stokesnetwork/stokes/releases/new\n"
echo -e "${YELLOW}2. Fill in the release form:${NC}"
echo -e "   - Tag: ${BLUE}$VERSION${NC}"
echo -e "   - Title: ${BLUE}Stokes $VERSION${NC}"
echo -e "   - Description: Copy from below\n"
echo -e "${YELLOW}3. Upload these files (drag & drop):${NC}"
for archive in "$RELEASE_DIR"/*.{tar.gz,zip} 2>/dev/null; do
    [ -f "$archive" ] && echo -e "   - $(basename "$archive")"
done

echo -e "\n${YELLOW}4. Mark as pre-release if testnet${NC}"
echo -e "\n${YELLOW}5. Click 'Publish release'${NC}\n"

echo -e "${GREEN}=== Release Description (copy this) ===${NC}\n"
cat << EOF
## Stokes $VERSION

### 🚀 Testnet Release

This is a **testnet release** for public testing. Do not use for production.

### 📦 Downloads

Choose the appropriate binary for your system:

- **Linux (x64)**: \`stokes-$VERSION-linux-amd64.tar.gz\`
- **Linux (ARM64)**: \`stokes-$VERSION-linux-arm64.tar.gz\`
- **macOS (Intel)**: \`stokes-$VERSION-macos-amd64.tar.gz\`
- **macOS (Apple Silicon)**: \`stokes-$VERSION-macos-arm64.tar.gz\`
- **Windows (x64)**: \`stokes-$VERSION-windows-amd64.zip\`

### 📚 Quick Start

1. Download and extract the archive for your platform
2. Follow the [Quick Start Guide](https://github.com/stokesnetwork/stokes#-quick-start-testnet)
3. Join our community and report any issues!

### ⚠️ Important Notes

- This is **testnet only** - coins have no real value
- Mainnet launch date: TBA
- Report bugs: [GitHub Issues](https://github.com/stokesnetwork/stokes/issues)

### 🔗 Links

- **Website**: https://stokesnetwork.github.io/stokes
- **Documentation**: [README.md](https://github.com/stokesnetwork/stokes/blob/master/README.md)
- **Twitter**: [@StokesCoin](https://twitter.com/StokesCoin)
EOF

echo -e "\n${GREEN}=== End of Description ===${NC}\n"
