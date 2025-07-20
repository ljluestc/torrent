#!/bin/bash
set -e

echo "===================================================================="
echo "       Temporarily disabling problematic directories"
echo "===================================================================="

# Create a backup of the current state
BACKUP_DIR="torrent_backup_$(date +%Y%m%d_%H%M%S)"
echo "Creating backup in $BACKUP_DIR..."
mkdir -p "$BACKUP_DIR"
cp -r . "$BACKUP_DIR"

echo "Temporarily renaming problematic directories..."

# Rename fs directory
if [ -d "fs" ]; then
    mv fs fs.disabled
    echo "  Renamed fs to fs.disabled"
fi

# Rename possum directory
if [ -d "storage/possum" ]; then
    mkdir -p storage/possum.disabled
    mv storage/possum/* storage/possum.disabled/
    echo "  Moved contents of storage/possum to storage/possum.disabled"
fi

echo "Now you can run tests on the rest of the codebase:"
echo "  go test ./..."

echo ""
echo "To restore the directories when you're done:"
echo "  mv fs.disabled fs"
echo "  mkdir -p storage/possum"
echo "  mv storage/possum.disabled/* storage/possum/"
echo "  rmdir storage/possum.disabled"
