#!/bin/bash
set -e

echo "===================================================================="
echo "               Cleaning up incorrectly placed files"
echo "===================================================================="

# Remove any incorrectly placed files
if [ -f "fix_test_helper_torrentfs.go" ]; then
    echo "Removing incorrect fix_test_helper_torrentfs.go from root directory..."
    rm fix_test_helper_torrentfs.go
fi

# Remove any other helper or fix scripts that might be causing issues
for file in fix_*.go; do
    if [ -f "$file" ]; then
        echo "Removing $file from root directory..."
        rm "$file"
    fi
done

echo "Cleanup complete."
