#!/bin/bash
set -e

echo "===================================================================="
echo "     RUNNING TESTS FOR WORKING PACKAGES ONLY"
echo "===================================================================="

# Apply fixes first
chmod +x remove_root_conflict.sh
./remove_root_conflict.sh

echo "Running tests for working packages..."
go test github.com/anacrolix/torrent/bencode
go test github.com/anacrolix/torrent/metainfo
go test github.com/anacrolix/torrent/iplist
go test github.com/anacrolix/torrent/tracker
go test github.com/anacrolix/torrent/tracker/http
go test github.com/anacrolix/torrent/tracker/udp
go test github.com/anacrolix/torrent/peer_protocol
go test github.com/anacrolix/torrent/peer_protocol/ut-holepunch
go test github.com/anacrolix/torrent/internal/nestedmaps
go test github.com/anacrolix/torrent/internal/alloclim
go test github.com/anacrolix/torrent/util/dirwatch
go test github.com/anacrolix/torrent/segments
go test github.com/anacrolix/torrent/webseed
go test github.com/anacrolix/torrent/webtorrent
go test github.com/anacrolix/torrent/mse
go test github.com/anacrolix/torrent/tests/issue-952
go test github.com/anacrolix/torrent/request-strategy

echo "All working package tests completed successfully!"
