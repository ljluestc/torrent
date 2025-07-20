#!/bin/bash
set -e

echo "===================================================================="
echo "       CLEANING ONLY ROOT DIRECTORY FILES"
echo "===================================================================="

# Find and list all Go files in the root directory
echo "Go files in root directory:"
ls -la *.go

# Check for files with package torrentfs in the root
ROOT_FILES_WITH_TORRENTFS=$(grep -l "package torrentfs" *.go 2>/dev/null || echo "")
if [ -n "$ROOT_FILES_WITH_TORRENTFS" ]; then
    echo "Found files with package torrentfs in root:"
    echo "$ROOT_FILES_WITH_TORRENTFS"
    
    # Create a backup of these files
    BACKUP_DIR="root_files_backup_$(date +%Y%m%d_%H%M%S)"
    mkdir -p "$BACKUP_DIR"
    
    for file in $ROOT_FILES_WITH_TORRENTFS; do
        cp "$file" "$BACKUP_DIR/"
        rm "$file"
        echo "  Removed $file (backup in $BACKUP_DIR)"
    done
else
    echo "No files with package torrentfs found in root directory."
fi

# Run a basic test of selected packages
echo "Running tests for a few selected packages..."
go test github.com/anacrolix/torrent/bencode
go test github.com/anacrolix/torrent/metainfo
go test github.com/anacrolix/torrent/iplist

echo "===================================================================="
echo "                   ROOT DIRECTORY CLEANED"
echo "===================================================================="
