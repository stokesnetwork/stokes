#!/bin/bash
# Script to create a new release tag for Stokes

set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== Stokes Release Creator ===${NC}\n"

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

echo -e "${YELLOW}Creating release: $VERSION${NC}\n"

# Check if tag already exists
if git rev-parse "$VERSION" >/dev/null 2>&1; then
    echo -e "${RED}Error: Tag $VERSION already exists${NC}"
    exit 1
fi

# Check if working directory is clean
if [ -n "$(git status --porcelain)" ]; then
    echo -e "${RED}Error: Working directory is not clean${NC}"
    echo -e "Please commit or stash your changes first"
    git status --short
    exit 1
fi

# Confirm with user
echo -e "${YELLOW}This will:${NC}"
echo "  1. Create git tag: $VERSION"
echo "  2. Push tag to origin"
echo "  3. Trigger GitHub Actions to build binaries"
echo "  4. Create a GitHub Release with all platform binaries"
echo ""
read -p "Continue? (y/N): " -n 1 -r
echo ""

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${RED}Aborted${NC}"
    exit 1
fi

# Create annotated tag
echo -e "\n${GREEN}Creating tag...${NC}"
git tag -a "$VERSION" -m "Release $VERSION"

# Push tag
echo -e "${GREEN}Pushing tag to origin...${NC}"
git push origin "$VERSION"

echo -e "\n${GREEN}✓ Success!${NC}"
echo -e "\nRelease tag $VERSION has been created and pushed."
echo -e "GitHub Actions is now building binaries for all platforms."
echo -e "\nCheck progress at:"
echo -e "  ${YELLOW}https://github.com/stokesnetwork/stokes/actions${NC}"
echo -e "\nRelease will be available at:"
echo -e "  ${YELLOW}https://github.com/stokesnetwork/stokes/releases/tag/$VERSION${NC}"
echo -e "\n${YELLOW}Note: It may take 10-15 minutes for all binaries to build.${NC}"
