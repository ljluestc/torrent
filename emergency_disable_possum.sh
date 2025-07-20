#!/bin/bash
set -e

echo "===================================================================="
echo "      EMERGENCY DISABLING OF POSSUM PACKAGE"
echo "===================================================================="

# Create a backup of the possum directory
BACKUP_DIR="possum_backup_$(date +%Y%m%d_%H%M%S)"
mkdir -p "storage/$BACKUP_DIR"
if [ -d "storage/possum" ]; then
    cp -r storage/possum/* "storage/$BACKUP_DIR"
    echo "Backup created in storage/$BACKUP_DIR"
fi

# Completely remove the possum directory and create a replacement
rm -rf storage/possum
mkdir -p storage/possum

# Create a minimal README to explain
cat > storage/possum/README.md << 'EOF'
# Possum Storage Package

This package has been temporarily disabled due to build issues.
The original code has been backed up.

Please run tests using:
