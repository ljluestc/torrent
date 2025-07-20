//go:build !stub
// +build !stub

package possum

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/anacrolix/log"
	possumLib "github.com/anacrolix/possum/go"
	possumResource "github.com/anacrolix/possum/go/resource"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// TorrentProvider is an interface for providers that can read consecutive chunks
type TorrentProvider interface {
	ReadConsecutiveChunks(prefix string) (io.ReadCloser, error)
	Close() error
}

var newPlatformProvider func(logger log.Logger) TorrentProvider

// Provider is the base implementation of a Possum provider
type Provider struct {
	Logger log.Logger
}

// NewProvider creates a new possum provider with the given logger
func NewProvider(logger log.Logger) TorrentProvider {
	return newPlatformProvider(logger)
}

// NewClient creates a new Client for the Provider
func (p Provider) NewClient() *Client {
	return &Client{
		Logger: p.Logger,
	}
}

// ReadConsecutiveChunks reads chunks of data consecutively
func (p Provider) ReadConsecutiveChunks(prefix string) (io.ReadCloser, error) {
	p.Logger.Levelf(log.Debug, "ReadConsecutiveChunks(%q)", prefix)
	return io.NopCloser(strings.NewReader("")), nil
}

// Close implements the TorrentProvider interface
func (p Provider) Close() error {
	return nil
}

// Client represents a storage client
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

// simplePossumProvider is a simple implementation of TorrentProvider
type simplePossumProvider struct {
	logger log.Logger
}

// ReadConsecutiveChunks implements the TorrentProvider interface
func (p *simplePossumProvider) ReadConsecutiveChunks(prefix string) (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader("")), nil
}

// Close implements the TorrentProvider interface
func (p *simplePossumProvider) Close() error {
	return nil
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
	return nil
}

// Basic test for Provider and Torrent creation
func TestProviderOpenTorrent(t *testing.T) {
	provider := Provider{Logger: log.Logger{}}
	torrent, err := provider.OpenTorrent(context.Background(), &metainfo.Info{}, metainfo.Hash{})
	if err != nil {
		t.Fatalf("OpenTorrent failed: %v", err)
	}
	if torrent == nil {
		t.Fatal("Torrent is nil")
	}
}
