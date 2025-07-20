#!/bin/bash
set -e

echo "===================================================================="
echo "     REMOVING CONFLICTING FILES IN ROOT DIRECTORY"
echo "===================================================================="

# Create a backup directory
BACKUP_DIR="conflicting_files_backup_$(date +%Y%m%d%H%M%S)"
mkdir -p "$BACKUP_DIR"

# Find any files with "package torrentfs" in the root directory
ROOT_FILES_WITH_TORRENTFS=$(grep -l "package torrentfs" *.go 2>/dev/null || echo "")

if [ -n "$ROOT_FILES_WITH_TORRENTFS" ]; then
    echo "Found files with package torrentfs in root directory:"
    echo "$ROOT_FILES_WITH_TORRENTFS"
    
    for file in $ROOT_FILES_WITH_TORRENTFS; do
        echo "Backing up and removing: $file"
        cp "$file" "$BACKUP_DIR/"
        rm "$file"
    done
    
    echo "All conflicting files have been backed up to $BACKUP_DIR and removed"
else
    echo "No conflicting files found in root directory"
fi

echo "Done cleaning root directory"
