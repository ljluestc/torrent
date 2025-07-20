#!/bin/bash
set -e

echo "===================================================================="
echo "                Building torrent package safely"
echo "===================================================================="

# First fix the package issues
echo "Fixing package issues..."
chmod +x fix_fs_package.sh fix_possum_package.sh
./fix_fs_package.sh
./fix_possum_package.sh

echo "Building project while excluding problematic packages..."
go build $(go list ./... | grep -v github.com/anacrolix/torrent/fs | grep -v github.com/anacrolix/torrent/storage/possum)

echo "===================================================================="
echo "                        Build complete"
echo "===================================================================="
echo ""
echo "To build with stub implementations, use:"
echo "  go build -tags stub ./..."
echo ""
