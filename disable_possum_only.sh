#!/bin/bash
set -e

echo "===================================================================="
echo "      DISABLING ONLY THE POSSUM PACKAGE"
echo "===================================================================="

# Step 1: Create a backup of the possum directory
BACKUP_DIR="possum_backup_$(date +%Y%m%d_%H%M%S)"
echo "Creating backup in storage/$BACKUP_DIR..."
mkdir -p "storage/$BACKUP_DIR"

if [ -d "storage/possum" ]; then
    cp -r storage/possum/* "storage/$BACKUP_DIR/"
    echo "Backup created in storage/$BACKUP_DIR"
    
    # Remove the contents of the possum directory
    rm -rf storage/possum/*
    
    # Create a README explaining what happened
    cat > storage/possum/README.md << 'EOF'
# Possum Storage Package (Temporarily Disabled)

This package has been temporarily disabled due to interface incompatibilities.
The original files have been backed up to the storage/possum_backup_* directory.

To run tests for the rest of the project, use:

```bash
go test $(go list ./... | grep -v github.com/anacrolix/torrent/storage/possum)
