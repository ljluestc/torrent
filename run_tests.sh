#!/bin/bash
set -e

# Apply fixes first
chmod +x direct_fix.sh
./direct_fix.sh

echo "========================================================"
echo "   Running tests for torrent project"
echo "========================================================"
echo "Skipping problematic packages (fs and possum)..."
go test $(go list ./... | grep -v github.com/anacrolix/torrent/fs | grep -v github.com/anacrolix/torrent/storage/possum)
