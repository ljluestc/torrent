//go:build stub
// +build stub
//go:build stub
// +build stub

package possum

import (
	"context"
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

// Client is a stub implementation of a storage client
type Client struct {
	Logger log.Logger
}

// Close implements client closer for the stub build
func (c *Client) Close() error {
	return nil
}

// OpenTorrent opens a torrent for storage (stub implementation)
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
	}, nil
}

// Close closes the provider (stub implementation)
func (p Provider) Close() error {
	return nil
}
// This file is used when the "stub" build tag is specified
package possum

import (
	"context"
	"errors"
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// OpenTorrent creates a stub implementation of a Torrent
func (p Provider) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (*Torrent, error) {
	return &Torrent{
		Logger: p.Logger,
	}, errors.New("possum stub: operation not supported")
}

// NewClient creates a new stub Client
func (p Provider) NewClient() *Client {
	return &Client{
		Logger: p.Logger,
	}
}

// GetIsComplete implements the interface for stub implementation
func (p *Piece) GetIsComplete() bool {
	return false
}
//go:build stub
// +build stub

// This file is used for stub implementation (when the "stub" build tag is specified)
package possum

import (
	"context"
	
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// Stub Provider implementation
type Provider struct {
	Logger log.Logger
}

// Stub Torrent implementation
type Torrent struct {
	Logger log.Logger
}

// NewClient creates a new client stub
func (p Provider) NewClient() *Client {
	return &Client{
		Logger: p.Logger,
	}
}

// Piece returns a stub piece implementation
func (t *Torrent) Piece(p metainfo.Piece) shared.PieceImpl {
	return &Piece{}
}

// Close is a no-op for the stub
func (t *Torrent) Close() error {
	return nil
}

// Stub Client implementation
type Client struct {
	Logger log.Logger
}

// Close is a no-op for the stub
func (c *Client) Close() error {
	return nil
}

// OpenTorrent returns a stub torrent
func (c *Client) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (shared.TorrentImpl, error) {
	return &Torrent{Logger: c.Logger}, nil
}

// Ensure interfaces are implemented
var (
	_ shared.TorrentImpl = &Torrent{}
	_ shared.PieceImpl = &Piece{}
)
// Ensure interface compliance for stub implementation
var (
	_ interface{ Close() error } = &Client{}
	_ shared.TorrentImpl         = &Torrent{}
	_ shared.PieceImpl           = &Piece{}
)
