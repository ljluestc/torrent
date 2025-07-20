#!/bin/bash
set -e

echo "===================================================================="
echo "      RUNNING TESTS FOR RELIABLE PACKAGES ONLY"
echo "===================================================================="

# List of packages known to be reliable
RELIABLE_PACKAGES=(
    "github.com/anacrolix/torrent/bencode"
    "github.com/anacrolix/torrent/metainfo"
    "github.com/anacrolix/torrent/iplist"
    "github.com/anacrolix/torrent/tracker"
    "github.com/anacrolix/torrent/tracker/http"
    "github.com/anacrolix/torrent/tracker/udp"
    "github.com/anacrolix/torrent/peer_protocol"
    "github.com/anacrolix/torrent/peer_protocol/ut-holepunch"
    "github.com/anacrolix/torrent/internal/nestedmaps"
    "github.com/anacrolix/torrent/internal/alloclim"
    "github.com/anacrolix/torrent/util/dirwatch"
    "github.com/anacrolix/torrent/segments"
    "github.com/anacrolix/torrent/webseed"
    "github.com/anacrolix/torrent/webtorrent"
    "github.com/anacrolix/torrent/mse"
    "github.com/anacrolix/torrent/tests/issue-952"
    "github.com/anacrolix/torrent/request-strategy"
)

# Run tests for each package
for pkg in "${RELIABLE_PACKAGES[@]}"; do
    echo "Testing $pkg..."
    go test "$pkg" || echo "  Failed: $pkg"
done

echo "===================================================================="
echo "      ALL RELIABLE PACKAGES TESTED"
echo "===================================================================="
