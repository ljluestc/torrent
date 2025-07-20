#!/bin/bash
set -e

echo "Fixing fs package declarations..."

# Find all Go files in the fs directory that declare package fs
files=$(grep -l "^package fs$" fs/*.go 2>/dev/null || true)

if [ -z "$files" ]; then
  echo "No files with 'package fs' declaration found."
else
  echo "Found files with 'package fs' declaration:"
  echo "$files"
  
  # Replace package declarations
  for file in $files; do
    echo "Fixing $file..."
    sed -i 's/^package fs$/package torrentfs/' "$file"
  done
  
  echo "All files fixed."
fi

# Check if there are still mixed package declarations
mixed=$(find fs -name "*.go" -exec grep -l "^package" {} \; | xargs grep -l "^package" | sort | uniq)
count=$(echo "$mixed" | wc -l)

if [ "$count" -gt 1 ]; then
  echo "WARNING: Still found mixed package declarations in fs directory:"
  for f in $mixed; do
    pkg=$(grep "^package" "$f" | head -n 1)
    echo "$f: $pkg"
  done
else
  echo "All files in fs directory now use the same package declaration."
fi
