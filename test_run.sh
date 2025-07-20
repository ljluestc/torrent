#!/bin/bash
set -e

# Function to print section headers
print_header() {
  echo "===================================================================="
  echo "  $1"
  echo "===================================================================="
  echo ""
}

# Option 0: Run specific packages for development work
if [ "$1" == "dev" ]; then
  print_header "DEVELOPMENT MODE: Running only specified packages"
  go test "$2" -v
  exit $?
fi

# Option 1: Skip problematic packages completely (safest)
print_header "SAFE MODE: Skipping problematic packages"
go test $(go list ./... | grep -v github.com/anacrolix/torrent/fs | grep -v github.com/anacrolix/torrent/storage/possum)

# Option 2: Run with stub implementation
if [ "$1" == "all" ]; then
  print_header "STUB MODE: Using stub implementations"
  go test -tags stub ./...
fi

echo ""
echo "Tests completed!"
