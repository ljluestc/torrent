#!/bin/bash
set -e

echo "===================================================================="
echo "      FIXING ROOT DIRECTORY PACKAGE CONFLICT"
echo "===================================================================="

# Create a backup of the conflicting files
BACKUP_DIR="root_backup_$(date +%Y%m%d_%H%M%S)"
mkdir -p "$BACKUP_DIR"

# Find any files with package torrentfs in the root directory
echo "Finding files with package torrentfs in root directory..."
ROOT_FILES_WITH_TORRENTFS=$(grep -l "package torrentfs" *.go 2>/dev/null || echo "")
if [ -n "$ROOT_FILES_WITH_TORRENTFS" ]; then
    echo "Found files with package torrentfs in root:"
    echo "$ROOT_FILES_WITH_TORRENTFS"
    
    # Backup the files before removing them
    for file in $ROOT_FILES_WITH_TORRENTFS; do
        echo "  Backing up $file to $BACKUP_DIR"
        cp "$file" "$BACKUP_DIR/"
        
        echo "  Removing $file"
        rm "$file"
    done
    echo "All conflicting files have been backed up and removed."
else
    echo "No files with package torrentfs found in root directory."
fi

echo "===================================================================="
echo "      ROOT DIRECTORY PACKAGE CONFLICT FIXED"
echo "===================================================================="
echo ""
echo "You can now try running tests on working packages with:"
echo "  ./run_reliable_tests.sh"
echo ""
echo "Your original files are backed up in: $BACKUP_DIR"
