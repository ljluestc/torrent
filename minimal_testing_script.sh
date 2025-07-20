#!/bin/bash
set -e

echo "===================================================================="
echo "             Running minimal set of working tests"
echo "===================================================================="

# Run tests for specific working packages
echo "Running tests for known working packages..."

echo "Testing bencode package..."
go test github.com/anacrolix/torrent/bencode

echo "Testing metainfo package..."
go test github.com/anacrolix/torrent/metainfo

echo "Testing iplist package..."
go test github.com/anacrolix/torrent/iplist

echo "Testing tracker packages..."
go test github.com/anacrolix/torrent/tracker
go test github.com/anacrolix/torrent/tracker/http
go test github.com/anacrolix/torrent/tracker/udp

echo "Testing peer_protocol packages..."
go test github.com/anacrolix/torrent/peer_protocol
go test github.com/anacrolix/torrent/peer_protocol/ut-holepunch

echo "Testing additional packages..."
go test github.com/anacrolix/torrent/internal/nestedmaps
go test github.com/anacrolix/torrent/internal/alloclim
go test github.com/anacrolix/torrent/util/dirwatch
go test github.com/anacrolix/torrent/segments
go test github.com/anacrolix/torrent/webseed
go test github.com/anacrolix/torrent/webtorrent
go test github.com/anacrolix/torrent/mse

echo "===================================================================="
echo "             All specified tests completed successfully!"
echo "===================================================================="
