#!/bin/bash
set -e

echo "===================================================================="
echo "      QUICK FIX AND TEST"
echo "===================================================================="

# First, fix the root directory
./fix_root_directory.sh

# Then, run tests on reliable packages
./run_reliable_tests.sh

echo "===================================================================="
echo "      QUICK FIX AND TEST COMPLETED"
echo "===================================================================="
