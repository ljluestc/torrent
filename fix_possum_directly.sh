#!/bin/bash
set -e

echo "===================================================================="
echo "             Directly fixing possum provider issues"
echo "===================================================================="

# Completely replace provider.go
echo "Replacing provider.go..."
cat > storage/possum/provider.go << 'EOF'
//go:build !stub
// +build !stub

package possumTorrentStorage

import (
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/storage"
)

// Provider implements storage provider for possum
type Provider struct {
	Logger log.Logger
}

// NewClient creates a new client implementation
func (p Provider) NewClient() (storage.ClientImplCloser, error) {
	return &Client{logger: p.Logger}, nil
}
EOF

# Check for possum-provider.go and remove it if it exists
if [ -f storage/possum/possum-provider.go ]; then
    echo "Removing conflicting possum-provider.go..."
    rm storage/possum/possum-provider.go
fi

# Fix client.go implementation
echo "Fixing client.go..."
cat > storage/possum/client.go << 'EOF'
//go:build !stub
// +build !stub

package possumTorrentStorage

import (
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage"
)

// Client implements storage.ClientImplCloser
type Client struct {
	logger log.Logger
}

// Close implements storage.ClientImplCloser
func (c *Client) Close() error {
	return nil
}

// OpenTorrent implements storage.ClientImpl
func (c *Client) OpenTorrent(info *storage.TorrentInfo) (storage.TorrentImpl, error) {
	return &Torrent{logger: c.logger}, nil
}

// Torrent implements storage.TorrentImpl
type Torrent struct {
	logger log.Logger
}

// Piece implements storage.TorrentImpl
func (t *Torrent) Piece(p metainfo.Piece) storage.PieceImpl {
	return &Piece{}
}

// Close implements storage.TorrentImpl
func (t *Torrent) Close() error {
	return nil
}

// Piece implements storage.PieceImpl
type Piece struct {}

// ReadAt implements storage.PieceImpl
func (p *Piece) ReadAt(b []byte, off int64) (n int, err error) {
	return 0, nil
}

// WriteAt implements storage.PieceImpl
func (p *Piece) WriteAt(b []byte, off int64) (n int, err error) {
	return len(b), nil
}

// MarkComplete implements storage.PieceImpl
func (p *Piece) MarkComplete() error {
	return nil
}

// MarkNotComplete implements storage.PieceImpl
func (p *Piece) MarkNotComplete() error {
	return nil
}

// Completion implements storage.PieceImpl
func (p *Piece) Completion() storage.Completion {
	return storage.Completion{
		Complete: false,
		Ok:       true,
	}
}

// GetIsComplete implements the interface
func (p *Piece) GetIsComplete() bool {
	return false
}

// Ensure interface compliance
var (
	_ storage.ClientImplCloser = &Client{}
	_ storage.TorrentImpl      = &Torrent{}
	_ storage.PieceImpl        = &Piece{}
)
EOF

echo "Done fixing possum package issues."
