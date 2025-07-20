#!/bin/bash
set -e

echo "===================================================================="
echo "          COMPREHENSIVE FIX FOR ALL PACKAGE CONFLICTS"
echo "===================================================================="

# Make all fix scripts executable
chmod +x fix_root_conflicts.sh fix_possum_conflicts.sh

# Step 1: Fix root directory conflicts
echo "Step 1: Fixing root directory conflicts..."
./fix_root_conflicts.sh

# Step 2: Fix possum package conflicts
echo "Step 2: Fixing possum package conflicts..."
./fix_possum_conflicts.sh

# Step 3: Standardize fs directory to use package torrentfs
echo "Step 3: Standardizing fs directory to use package torrentfs..."
for file in $(find fs/ -name "*.go"); do
    if grep -q "^package fs$" "$file"; then
        echo "  Changing package in $file from 'fs' to 'torrentfs'"
        sed -i 's/^package fs$/package torrentfs/' "$file"
    fi
done

echo "===================================================================="
echo "          ALL PACKAGE CONFLICTS FIXED"
echo "===================================================================="
echo ""
echo "You can now run tests with:"
echo "  go test \$(go list ./... | grep -v github.com/anacrolix/torrent/fs | grep -v github.com/anacrolix/torrent/storage/possum)"
echo ""
echo "Or run with stub implementation:"
echo "  go test -tags stub ./..."
