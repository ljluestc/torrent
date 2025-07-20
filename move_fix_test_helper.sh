#!/bin/bash
set -e

echo "===================================================================="
echo "      MOVING PROBLEMATIC TEST HELPER FILE"
echo "===================================================================="

# Create a backup of the file first
BACKUP_DIR="backup_$(date +%Y%m%d_%H%M%S)"
mkdir -p "$BACKUP_DIR"

if [ -f "fix_test_helper_torrentfs.go" ]; then
    echo "Found problematic file: fix_test_helper_torrentfs.go"
    cp "fix_test_helper_torrentfs.go" "$BACKUP_DIR/"
    echo "Backed up file to $BACKUP_DIR/fix_test_helper_torrentfs.go"
    
    # Move the file to the fs directory where it belongs
    echo "Moving file to fs directory..."
    mkdir -p fs
    mv fix_test_helper_torrentfs.go fs/
    echo "File moved successfully"
else
    echo "Could not find fix_test_helper_torrentfs.go in root directory"
    
    # Look for any file with package torrentfs in the root
    FILES_WITH_TORRENTFS=$(grep -l "package torrentfs" *.go 2>/dev/null || echo "")
    if [ -n "$FILES_WITH_TORRENTFS" ]; then
        echo "Found other files with package torrentfs in root:"
        echo "$FILES_WITH_TORRENTFS"
        
        for file in $FILES_WITH_TORRENTFS; do
            echo "Backing up $file"
            cp "$file" "$BACKUP_DIR/"
            
            echo "Moving $file to fs directory"
            mkdir -p fs
            mv "$file" fs/
        done
        echo "All problematic files moved to fs directory"
    else
        echo "No files with package torrentfs found in root directory"
    fi
fi

echo "===================================================================="
echo "      FILE MOVE OPERATION COMPLETE"
echo "===================================================================="
echo ""
echo "You can now try running tests on working packages with:"
echo "  ./run_tests_on_working.sh"
