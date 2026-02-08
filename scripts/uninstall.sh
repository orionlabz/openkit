#!/usr/bin/env bash
set -e

# OpenKit Uninstall Script
# This script removes OpenKit CLI from your system
#
# Usage:
#   bash uninstall.sh           # Interactive mode (asks for confirmation)
#   bash uninstall.sh -y        # Non-interactive mode (auto-confirm)
#   curl ... | bash             # Piped mode (auto-confirm, no stdin available)

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Parse arguments
AUTO_CONFIRM=false
if [[ "$1" == "-y" ]] || [[ "$1" == "--yes" ]]; then
    AUTO_CONFIRM=true
fi

# Check if running in pipe mode (stdin not available)
if [ ! -t 0 ]; then
    AUTO_CONFIRM=true
fi

echo -e "${GREEN}OpenKit Uninstall Script${NC}"
echo ""

# Check if openkit is installed
if ! command -v openkit &> /dev/null; then
    echo -e "${YELLOW}OpenKit is not installed on this system.${NC}"
    exit 0
fi

# Get installation location
INSTALL_PATH=$(which openkit)
echo -e "Found OpenKit at: ${YELLOW}${INSTALL_PATH}${NC}"

# Get current version
CURRENT_VERSION=$(openkit version 2>/dev/null | grep "Version:" | awk '{print $2}' || echo "unknown")
echo -e "Current version: ${YELLOW}${CURRENT_VERSION}${NC}"
echo ""

# Confirm removal
if [ "$AUTO_CONFIRM" = false ]; then
    read -p "Do you want to remove OpenKit? [y/N] " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo -e "${YELLOW}Uninstallation cancelled.${NC}"
        exit 0
    fi
else
    echo -e "${YELLOW}Auto-confirming removal (non-interactive mode)${NC}"
fi

# Remove binary
echo -e "Removing OpenKit binary..."
if [ -w "$INSTALL_PATH" ]; then
    rm "$INSTALL_PATH"
else
    sudo rm "$INSTALL_PATH"
fi

# Verify removal
if command -v openkit &> /dev/null; then
    echo -e "${RED}Failed to remove OpenKit.${NC}"
    exit 1
else
    echo -e "${GREEN}OpenKit successfully removed!${NC}"
fi

echo ""
echo -e "${GREEN}Uninstallation complete.${NC}"
echo ""
echo "Note: This script does not remove:"
echo "  - Project-specific OpenKit configurations (.openkit/, .opencode/, etc.)"
echo "  - Go installation (if used)"
echo ""
echo "To reinstall OpenKit, visit:"
echo "  https://github.com/openkit-devtools/openkit"
