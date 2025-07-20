#!/bin/bash

echo "===================================================================="
echo "             Running all tests except fs packages"
echo "===================================================================="

echo "Running tests for all packages except fs-related..."
go test $(go list ./... | grep -v github.com/anacrolix/torrent/fs)

echo "===================================================================="
echo "                        Tests complete"
echo "===================================================================="
