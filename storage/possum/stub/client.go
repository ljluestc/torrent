// Package stub provides stub implementations for possum
package stub

import (
	"context"
	"errors"

	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// Client is a stub implementation for the possum storage client
type Client struct {
	Logger log.Logger
}

// Close implements storage.ClientImplCloser
func (c *Client) Close() error {
	return nil
}

// OpenTorrent returns a stub torrent implementation
func (c *Client) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (shared.TorrentImpl, error) {
	return &StubTorrent{}, errors.New("stub implementation - not supported")
}

// StubTorrent is a stub implementation of shared.TorrentImpl
type StubTorrent struct{}

// Piece returns a stub piece implementation
func (t *StubTorrent) Piece(p metainfo.Piece) shared.PieceImpl {
	return &Piece{}
}

// Close closes the stub torrent
func (t *StubTorrent) Close() error {
	return nil
}
