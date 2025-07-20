#!/bin/bash
set -e

echo "Fixing build issues..."

# Fix package conflicts
find fs -name "*.go" -exec sed -i 's/^package fs$/package torrentfs/g' {} \;

# Make sure stub.c is removed
rm -f storage/possum/stub.c

# Add build tags to all possum files that don't have them
find storage/possum -name "*.go" | grep -v "_test.go" | grep -v "stub" | xargs -I{} bash -c '
  if ! grep -q "//go:build" {}; then
    sed -i "1s/^/\/\/go:build !stub\n\/\/ +build !stub\n\n/" {}
  fi
'

echo "Build fixes applied successfully!"
