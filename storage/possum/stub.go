//go:build stub
// +build stub

package possum

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// Provider implements a stub provider for the stub build
type Provider struct {
	Logger log.Logger
}

// Stub Torrent implementation
type Torrent struct {
	Logger log.Logger
}

// Client is a stub implementation of a storage client
type Client struct {
	Logger log.Logger
}

// Close implements client closer for the stub build
func (c *Client) Close() error {
	return nil
}

// OpenTorrent returns a stub torrent
func (c *Client) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (shared.TorrentImpl, error) {
	return &Torrent{Logger: c.Logger}, nil
}

// NewClient creates a new Client for the Provider (stub implementation)
func (p Provider) NewClient() *Client {
	return &Client{
		Logger: p.Logger,
	}
}

// ReadConsecutiveChunks implements the stub version
func (p Provider) ReadConsecutiveChunks(prefix string) (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader("")), nil
}

// OpenTorrent creates a new Torrent instance (stub implementation)
func (p Provider) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (*Torrent, error) {
	return &Torrent{
		Logger: p.Logger,
	}, errors.New("possum stub: operation not supported")
}

// Close closes the provider (stub implementation)
func (p Provider) Close() error {
	return nil
}

// Piece returns a stub piece implementation
func (t *Torrent) Piece(p metainfo.Piece) shared.PieceImpl {
	return &Piece{}
}

// GetIsComplete implements the interface for stub implementation
func (p *Piece) GetIsComplete() bool {
	return false
}

// Piece implements shared.PieceImpl
type Piece struct{}

// Ensure interfaces are implemented
var (
	_ interface{ Close() error } = &Client{}
	_ shared.TorrentImpl         = &Torrent{}
	_ shared.PieceImpl           = &Piece{}
)
