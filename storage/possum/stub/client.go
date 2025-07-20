//go:build stub
// +build stub

package stub

import (
	"context"
	"errors"
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
	"github.com/anacrolix/torrent/storage/possum"
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
	return &possum.Torrent{
		Logger: c.Logger,
	}, errors.New("stub implementation - not supported")
}
