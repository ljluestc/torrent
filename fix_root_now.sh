#!/bin/bash
set -e

echo "===================================================================="
echo "          FIXING ROOT DIRECTORY PACKAGE CONFLICTS"
echo "===================================================================="

# Find files with package torrentfs
echo "Finding files with package torrentfs in root directory..."
TORRENTFS_FILES=$(grep -l "package torrentfs" *.go 2>/dev/null || echo "")

if [ -n "$TORRENTFS_FILES" ]; then
    echo "Found files with package torrentfs in root directory:"
    echo "$TORRENTFS_FILES"
    
    echo "Removing these files..."
    for file in $TORRENTFS_FILES; do
        rm -f "$file"
        echo "  Removed $file"
    done
else
    echo "No files with package torrentfs found in root directory."
fi

# Look for any other package conflicts
echo "Checking for other package conflicts..."

# Find any Go files with package different from torrent
NONSTANDARD_PKG_FILES=$(grep -l "^package " --include="*.go" . | xargs grep -l "^package " --include="*.go" | grep -v "package torrent" 2>/dev/null || echo "")

if [ -n "$NONSTANDARD_PKG_FILES" ]; then
    echo "Found additional files with non-standard packages:"
    echo "$NONSTANDARD_PKG_FILES"
    
    echo "Examining these files..."
    for file in $NONSTANDARD_PKG_FILES; do
        # Only process files in the root directory
        if [[ "$file" != *"/"* ]]; then
            PKG_NAME=$(grep "^package " "$file" | head -1 | awk '{print $2}')
            echo "  $file has package $PKG_NAME"
            if [ "$PKG_NAME" != "torrent" ] && [ "$PKG_NAME" != "main" ]; then
                echo "  Removing conflicting file $file"
                rm -f "$file"
            fi
        fi
    done
else
    echo "No additional package conflicts found."
fi

echo "Root directory package conflicts fixed."
echo "===================================================================="
