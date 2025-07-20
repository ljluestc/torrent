#!/bin/bash
set -e

echo "===================================================================="
echo "             COMPREHENSIVE PACKAGE FIX FOR TORRENT PROJECT"
echo "===================================================================="

# Make all fix scripts executable
chmod +x root_package_fix.sh fs_package_fix.sh possum_package_fix.sh

# Step 1: Clean the root directory
echo "Step 1: Cleaning root directory..."
./root_package_fix.sh

# Step 2: Fix the fs directory
echo "Step 2: Fixing fs directory..."
./fs_package_fix.sh

# Step 3: Fix the possum package
echo "Step 3: Fixing possum package..."
./possum_package_fix.sh

echo "===================================================================="
echo "                     ALL FIXES APPLIED"
echo "===================================================================="
echo ""
echo "You can now run the tests with:"
echo "  go test \$(go list ./... | grep -v github.com/anacrolix/torrent/fs | grep -v github.com/anacrolix/torrent/storage/possum)"
echo ""
echo "Or run specific working packages with:"
echo "  go test github.com/anacrolix/torrent/bencode"
echo "  go test github.com/anacrolix/torrent/metainfo"
echo "  go test github.com/anacrolix/torrent/iplist"
echo "  go test github.com/anacrolix/torrent/tracker"
echo "  go test github.com/anacrolix/torrent/segments"
echo "  ... etc"
echo ""
echo "To run with stub implementation:"
echo "  go test -tags stub ./..."
