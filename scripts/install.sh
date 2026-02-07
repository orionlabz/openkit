#!/bin/bash
set -e

# OpenKit CLI Installer for macOS/Linux/WSL
# Usage: curl -fsSL https://openkit.dev/install | bash

REPO="openkit-dev/cli"
BINARY_NAME="openkit"
INSTALL_DIR="/usr/local/bin"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${CYAN}"
echo "   ___                   _  ___ _   "
echo "  / _ \ _ __   ___ _ __ | |/ (_) |_ "
echo " | | | | '_ \ / _ \ '_ \| ' /| | __|"
echo " | |_| | |_) |  __/ | | | . \| | |_ "
echo "  \___/| .__/ \___|_| |_|_|\_\_|\__|"
echo "       |_|                          "
echo -e "${NC}"
echo "Universal Spec-Driven Development Toolkit"
echo ""

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
  linux)
    OS="Linux"
    ;;
  darwin)
    OS="Darwin"
    ;;
  *)
    echo -e "${RED}Unsupported OS: $OS${NC}"
    exit 1
    ;;
esac

case "$ARCH" in
  x86_64)
    ARCH="x86_64"
    ;;
  aarch64|arm64)
    ARCH="arm64"
    ;;
  *)
    echo -e "${RED}Unsupported architecture: $ARCH${NC}"
    exit 1
    ;;
esac

echo -e "${CYAN}Detecting system...${NC}"
echo "  OS: $OS"
echo "  Architecture: $ARCH"
echo ""

# Get latest release
echo -e "${CYAN}Fetching latest release...${NC}"
LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_RELEASE" ]; then
  echo -e "${RED}Failed to fetch latest release${NC}"
  exit 1
fi

echo "  Latest version: $LATEST_RELEASE"
echo ""

# Construct download URL
# GoReleaser uses "cli" as the project name, not "openkit"
FILENAME="cli_${OS}_${ARCH}.tar.gz"
if [ "$OS" = "windows" ]; then
  FILENAME="cli_${OS}_${ARCH}.zip"
fi

DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/$FILENAME"

echo -e "${CYAN}Downloading OpenKit CLI...${NC}"
echo "  URL: $DOWNLOAD_URL"

# Create temp directory
TMP_DIR=$(mktemp -d)
trap "rm -rf $TMP_DIR" EXIT

cd "$TMP_DIR"

# Download
if ! curl -fsSL "$DOWNLOAD_URL" -o "$FILENAME"; then
  echo -e "${RED}Failed to download $FILENAME${NC}"
  exit 1
fi

# Extract
echo -e "${CYAN}Extracting...${NC}"
tar -xzf "$FILENAME"

# Install
echo -e "${CYAN}Installing to $INSTALL_DIR...${NC}"

if [ ! -w "$INSTALL_DIR" ]; then
  echo -e "${YELLOW}Requesting sudo permissions...${NC}"
  sudo install -m 755 "$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
else
  install -m 755 "$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
fi

# Verify installation
if command -v "$BINARY_NAME" &> /dev/null; then
  echo ""
  echo -e "${GREEN}✓ OpenKit CLI installed successfully!${NC}"
  echo ""
  echo "Run 'openkit --help' to get started"
  echo ""
  echo "Examples:"
  echo "  openkit check              # Check system requirements"
  echo "  openkit init my-app        # Create new project"
  echo "  openkit init --ai claude   # Create project for Claude Code"
  echo ""
else
  echo -e "${RED}Installation failed. Binary not found in PATH${NC}"
  exit 1
fi
