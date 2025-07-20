#!/bin/bash
set -e

echo "===================================================================="
echo "          COMPLETELY FIXING FS PACKAGE"
echo "===================================================================="

# Step 1: Convert all files in fs directory to use package torrentfs
echo "Converting all files in fs directory to use package torrentfs..."
for file in $(find fs/ -name "*.go"); do
    # Check if file contains "package fs"
    if grep -q "package fs" "$file"; then
        echo "  Converting $file from 'package fs' to 'package torrentfs'"
        sed -i 's/package fs/package torrentfs/g' "$file"
    fi
done

echo "===================================================================="
echo "          FS PACKAGE COMPLETELY FIXED"
echo "===================================================================="
