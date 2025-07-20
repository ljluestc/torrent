#!/bin/bash

echo "===================================================================="
echo "        Checking for package declaration issues in Go files"
echo "===================================================================="

# Check for mixed package declarations in directories
echo "Checking for directories with mixed package declarations..."

# Create a temporary file to store results
TEMP_FILE=$(mktemp)

# Find all Go files and their package declarations
find . -name "*.go" | while read -r file; do
    dir=$(dirname "$file")
    pkg=$(grep -E "^package [a-zA-Z0-9_]+" "$file" | head -n 1 | sed 's/^package //')
    echo "$dir $pkg" >> "$TEMP_FILE"
done

# Find directories with multiple package declarations
echo "Directories with mixed package declarations:"
sort "$TEMP_FILE" | awk '{print $1, $2}' | uniq | awk '{print $1}' | uniq -c | grep -v "^ *1 " | awk '{print $2}' | while read -r dir; do
    echo "  $dir:"
    grep "^$dir " "$TEMP_FILE" | awk '{print $2}' | sort | uniq -c | awk '{printf "    %s (%d files)\n", $2, $1}'
done

echo ""
echo "Checking for build tag issues..."
# Find files with multiple build tags
find . -name "*.go" | xargs grep -l "//go:build" | while read -r file; do
    count=$(grep -c "//go:build" "$file")
    if [ "$count" -gt 1 ]; then
        echo "  $file: multiple //go:build comments ($count)"
    fi
done

# Clean up
rm "$TEMP_FILE"

echo "===================================================================="
echo "                      Check complete!"
echo "===================================================================="
