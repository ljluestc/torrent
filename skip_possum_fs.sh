#!/bin/bash
set -e

echo "===================================================================="
echo "      RUNNING TESTS SKIPPING PROBLEMATIC PACKAGES"
echo "===================================================================="

# Run tests skipping problematic packages
echo "Running tests for non-problematic packages..."
go test $(go list ./... | grep -v github.com/anacrolix/torrent/fs | grep -v github.com/anacrolix/torrent/storage/possum)

echo "===================================================================="
echo "                TESTS COMPLETED"
echo "===================================================================="
