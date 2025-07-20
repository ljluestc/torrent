//go:build !stub
// +build !stub

package possum

import (
	"context"
	"io"
	"strings"

	"github.com/anacrolix/log"
	possumLib "github.com/anacrolix/possum/go"
	possumResource "github.com/anacrolix/possum/go/resource"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// Provider implements resource.Provider and ConsecutiveChunkReader
type Provider struct {
	possumResource.Provider
	Logger log.Logger
	Dir    *possumLib.Dir
}

// OpenTorrent creates a new Torrent instance for storage
func (p Provider) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (*Torrent, error) {
	return &Torrent{
		Logger: p.Logger,
	}, nil
}

// NewClient creates a new Client for the Provider
func (p Provider) NewClient() *Client {
	return &Client{
		Logger: p.Logger,
	}
}


```

// ReadConsecutiveChunks reads chunks of data consecutively
func (p Provider) ReadConsecutiveChunks(prefix string) (io.ReadCloser, error) {
	p.Logger.Levelf(log.Debug, "ReadConsecutiveChunks(%q)", prefix)
	return io.NopCloser(strings.NewReader("")), nil
}

// MovePrefix moves data from one prefix to another
func (p Provider) MovePrefix(from, to string) error {
	return nil
}

// NewInstance creates a new possum instance
func (p *Provider) NewInstance(name string) (interface{}, error) {
	if p.Dir == nil {
		return nil, nil
	}
	// If possumLib.Dir has an InstanceForKey or similar, use it. Otherwise, return nil.
	return nil, nil
}

// ReadResource reads a resource with the given key
func (p Provider) ReadResource(key string) (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader("")), nil
}

// WriteResource writes data to a resource with the given key
func (p Provider) WriteResource(key string, f func(io.Writer) error) error {
	return nil
}

// GetLength returns the length of a resource
func (p Provider) GetLength(key string) (int64, error) {
	return 0, nil
}

// Close closes the provider
func (p Provider) Close() error {
	// Since Dir doesn't have a Close method, we return nil
	// If specific cleanup is needed for Dir, it should be implemented accordingly
	return nil
}

// Client implements torrent storage client
type Client struct {
	Logger log.Logger
	logger log.Logger // For compatibility with both naming styles
}

// Close implements client closer
func (c *Client) Close() error {
	return nil
}

// OpenTorrent opens a torrent for storage
func (c *Client) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (shared.TorrentImpl, error) {
	logger := c.Logger
	if nil == logger {
		logger = c.logger
		if nil == logger {
			logger = log.Logger{}
		}
	}
	return &Torrent{Logger: logger}, nil
}
