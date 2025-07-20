#!/bin/bash
set -e

echo "===================================================================="
echo "    RUNNING EXTREME FIX AND TESTING WORKING PACKAGES"
echo "===================================================================="

# Make the extreme fix script executable
chmod +x extreme_fix.sh

# Run the extreme fix script
./extreme_fix.sh

# Run the working tests script that was created by extreme_fix.sh
./run_working_tests.sh

echo "===================================================================="
echo "    EXTREME FIX AND TESTS COMPLETED SUCCESSFULLY"
echo "===================================================================="
echo ""
echo "If you need to restore your project to its original state, run:"
echo "  ./restore_project.sh"
echo ""
echo "Your original code is backed up in a directory with name format:"
echo "  ../torrent_backup_YYYYMMDD_HHMMSS"
