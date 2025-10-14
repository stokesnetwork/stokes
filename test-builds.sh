#!/bin/bash
# Test cross-compilation for all platforms before creating a release

set -e

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${YELLOW}=== Testing Cross-Compilation for All Platforms ===${NC}\n"

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo -e "Go version: ${GREEN}$GO_VERSION${NC}"
if [[ ! "$GO_VERSION" =~ ^1\.(2[45]|[3-9][0-9]) ]]; then
    echo -e "${RED}Warning: Go 1.24+ recommended. You have $GO_VERSION${NC}\n"
fi
echo ""

# Create build directory
BUILD_DIR="./test-builds"
rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR"

# Platforms to test
declare -a PLATFORMS=(
    "linux:amd64"
    "linux:arm64"
    "darwin:amd64"
    "darwin:arm64"
    "windows:amd64"
)

FAILED=0
SUCCEEDED=0

for platform in "${PLATFORMS[@]}"; do
    IFS=':' read -r GOOS GOARCH <<< "$platform"
    
    echo -e "${YELLOW}Building for $GOOS/$GOARCH...${NC}"
    
    OUTPUT_DIR="$BUILD_DIR/$GOOS-$GOARCH"
    mkdir -p "$OUTPUT_DIR"
    
    # Determine binary suffix
    SUFFIX=""
    if [ "$GOOS" = "windows" ]; then
        SUFFIX=".exe"
    fi
    
    # Try building each binary
    BINARIES=("stokesd" "stokesctl:./cmd/stokesctl" "stokesminer:./cmd/stokesminer" "stokeswallet:./cmd/stokeswallet")
    
    PLATFORM_FAILED=0
    for binary_info in "${BINARIES[@]}"; do
        IFS=':' read -r binary_name binary_path <<< "$binary_info"
        if [ -z "$binary_path" ]; then
            binary_path="."
        fi
        
        echo -n "  - Building $binary_name... "
        
        if GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build \
            -o "$OUTPUT_DIR/${binary_name}${SUFFIX}" \
            "$binary_path" 2>&1 | tee "$OUTPUT_DIR/${binary_name}.build.log" > /dev/null; then
            echo -e "${GREEN}✓${NC}"
        else
            echo -e "${RED}✗${NC}"
            echo -e "${RED}    Error log saved to: $OUTPUT_DIR/${binary_name}.build.log${NC}"
            PLATFORM_FAILED=1
        fi
    done
    
    if [ $PLATFORM_FAILED -eq 0 ]; then
        echo -e "${GREEN}✓ $GOOS/$GOARCH: SUCCESS${NC}\n"
        SUCCEEDED=$((SUCCEEDED + 1))
    else
        echo -e "${RED}✗ $GOOS/$GOARCH: FAILED${NC}\n"
        FAILED=$((FAILED + 1))
    fi
done

echo -e "\n${YELLOW}=== Build Summary ===${NC}"
echo -e "Succeeded: ${GREEN}$SUCCEEDED${NC}"
echo -e "Failed: ${RED}$FAILED${NC}"

if [ $FAILED -gt 0 ]; then
    echo -e "\n${RED}Some builds failed. Check logs in $BUILD_DIR/${NC}"
    echo -e "${YELLOW}Fix the errors before creating a release.${NC}"
    exit 1
else
    echo -e "\n${GREEN}All builds successful! Ready to create a release.${NC}"
    exit 0
fi
