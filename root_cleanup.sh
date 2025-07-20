#!/bin/bash
set -e

echo "===================================================================="
echo "       CLEANING ROOT DIRECTORY OF PROBLEMATIC FILES"
echo "===================================================================="

# Create a backup directory
BACKUP_DIR="root_backups_$(date +%Y%m%d_%H%M%S)"
mkdir -p "$BACKUP_DIR"
echo "Created backup directory: $BACKUP_DIR"

# Look for the specific problematic file
if [ -f "fix_test_helper_torrentfs.go" ]; then
    echo "Found problematic file: fix_test_helper_torrentfs.go"
    cp "fix_test_helper_torrentfs.go" "$BACKUP_DIR/"
    rm "fix_test_helper_torrentfs.go"
    echo "Removed problematic file (backup saved)"
fi

# Find any other files with package torrentfs in the root
ROOT_FILES_WITH_TORRENTFS=$(grep -l "package torrentfs" *.go 2>/dev/null || echo "")
if [ -n "$ROOT_FILES_WITH_TORRENTFS" ]; then
    echo "Found other files with package torrentfs in root:"
    echo "$ROOT_FILES_WITH_TORRENTFS"
    
    for file in $ROOT_FILES_WITH_TORRENTFS; do
        echo "Backing up and removing $file"
        cp "$file" "$BACKUP_DIR/"
        rm "$file"
    done
    echo "All problematic files removed from root (backups saved)"
fi

echo "===================================================================="
echo "                   ROOT DIRECTORY CLEANED"
echo "===================================================================="
