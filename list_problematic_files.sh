#!/bin/bash

echo "===================================================================="
echo "       DIAGNOSTIC: LISTING ALL PROBLEMATIC FILES"
echo "===================================================================="

# Find files with package torrentfs in the root directory
echo "Files with 'package torrentfs' in root directory:"
grep -l "package torrentfs" *.go 2>/dev/null || echo "None found"
echo ""

# Check for mixed package declarations in the fs directory
echo "Package declarations in fs directory:"
find fs -name "*.go" -exec grep -l "^package" {} \; | xargs grep "^package" | sort | uniq -c
echo ""

# Check for build tag issues in possum directory
echo "Files with multiple build tags in storage/possum:"
find storage/possum -name "*.go" | xargs grep -l "//go:build" | while read file; do
    count=$(grep -c "//go:build" "$file")
    if [ "$count" -gt 1 ]; then
        echo "  $file: $count build tag comments"
    fi
done
echo ""

# List all Go files in the root directory
echo "All Go files in root directory:"
ls -la *.go 2>/dev/null || echo "None found"
echo ""

echo "===================================================================="
echo "                   DIAGNOSTIC COMPLETE"
echo "===================================================================="
