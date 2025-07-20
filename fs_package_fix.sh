#!/bin/bash
set -e

echo "===================================================================="
echo "          STANDARDIZING FS DIRECTORY TO PACKAGE TORRENTFS"
echo "===================================================================="

# Standardize all files in fs directory to use package torrentfs
echo "Converting all files in fs directory to use package torrentfs..."
for file in $(find fs/ -name "*.go"); do
    if grep -q "^package fs$" "$file"; then
        echo "  Changing package in $file from 'fs' to 'torrentfs'"
        sed -i 's/^package fs$/package torrentfs/' "$file"
    fi
done

# Make sure file_handle.go imports the correct torrent package
if grep -q "github.com/anacrolix/torrent/torrent" fs/file_handle.go; then
    echo "Fixing incorrect import path in file_handle.go..."
    sed -i 's|github.com/anacrolix/torrent/torrent|github.com/anacrolix/torrent|' fs/file_handle.go
fi

echo "All files in fs directory now use 'package torrentfs'"
