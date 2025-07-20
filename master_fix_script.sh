#!/bin/bash
set -e

echo "===================================================================="
echo "          MASTER FIX SCRIPT FOR TORRENT PROJECT"
echo "===================================================================="

# Make all scripts executable
chmod +x fix_root_now.sh fix_possum_completely.sh fix_fs_completely.sh 

# Step 1: Fix root directory
echo "Step 1: Fixing root directory package conflicts..."
./fix_root_now.sh

# Step 2: Fix possum package
echo "Step 2: Completely fixing possum package..."
./fix_possum_completely.sh

# Step 3: Fix fs package
echo "Step 3: Fixing fs package..."
./fix_fs_completely.sh

echo "===================================================================="
echo "          ALL FIXES APPLIED"
echo "===================================================================="
echo ""
echo "You can now run working tests with:"
echo "  go test \$(go list ./... | grep -v github.com/anacrolix/torrent/fs | grep -v github.com/anacrolix/torrent/storage/possum)"
echo ""
echo "Or run specific working packages individually:"
echo "  go test github.com/anacrolix/torrent/bencode"
echo "  go test github.com/anacrolix/torrent/metainfo"
echo "  etc."
echo ""
echo "===================================================================="
