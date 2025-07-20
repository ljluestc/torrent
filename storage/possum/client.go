//go:build !stub
// +build !stub

package possum

import (
	"context"
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// Client implements ClientImplCloser from the storage package
type Client struct {
	Logger log.Logger
}

// Close implements the Close method
func (c *Client) Close() error {
	return nil
}

// OpenTorrent opens a torrent for storage
func (c *Client) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (shared.TorrentImpl, error) {
	return &Torrent{
		Logger: c.Logger,
	}, nil
}
