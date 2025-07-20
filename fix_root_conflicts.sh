#!/bin/bash
set -e

echo "===================================================================="
echo "          FIXING CONFLICTS IN ROOT DIRECTORY"
echo "===================================================================="

# Find and remove any conflicting files in the root directory
ROOT_FILES_WITH_TORRENTFS=$(grep -l "package torrentfs" *.go 2>/dev/null || echo "")
if [ -n "$ROOT_FILES_WITH_TORRENTFS" ]; then
    echo "Found files in root with package torrentfs:"
    echo "$ROOT_FILES_WITH_TORRENTFS"
    echo "Removing these files..."
    for file in $ROOT_FILES_WITH_TORRENTFS; do
        rm "$file"
        echo "  Removed $file"
    done
fi

# Fix the test helper file if it exists
if [ -f "fix_test_helper_torrentfs.go" ]; then
    echo "Found problematic file: fix_test_helper_torrentfs.go"
    echo "Removing it..."
    rm fix_test_helper_torrentfs.go
    echo "File removed."
fi

echo "Root directory cleaned of conflicting files."
echo "===================================================================="
