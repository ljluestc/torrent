#!/bin/bash
set -e

echo "===================================================================="
echo "          REMOVING CONFLICTING FILES IN ROOT DIRECTORY"
echo "===================================================================="

# Remove the problematic file in the root directory
if [ -f "fix_test_helper_torrentfs.go" ]; then
    echo "Found problematic file: fix_test_helper_torrentfs.go"
    echo "Removing it..."
    rm fix_test_helper_torrentfs.go
    echo "File removed."
fi

# Check for any other Go files in the root with package torrentfs
ROOT_FILES_WITH_TORRENTFS=$(grep -l "package torrentfs" *.go 2>/dev/null || echo "")
if [ -n "$ROOT_FILES_WITH_TORRENTFS" ]; then
    echo "Found other files in root with package torrentfs:"
    echo "$ROOT_FILES_WITH_TORRENTFS"
    echo "Removing these files..."
    for file in $ROOT_FILES_WITH_TORRENTFS; do
        rm "$file"
        echo "  Removed $file"
    done
fi

echo "Root directory cleaned of conflicting files."
