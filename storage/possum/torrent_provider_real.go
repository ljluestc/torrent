//go:build !android && !windows
// +build !android,!windows

package possum

import (
	"io"
	"strings"

	"github.com/anacrolix/log"
)

func init() {
	// Set the platform-specific provider function
	newPlatformProvider = func(logger log.Logger) TorrentProvider {
		// In a real implementation, this would create a real possum provider
		// For now, we'll return a simple implementation
		return &simplePossumProvider{
			logger: logger,
		}
	}
}

// simplePossumProvider is a simple implementation of TorrentProvider
type simplePossumProvider struct {
	logger log.Logger
}

// ReadConsecutiveChunks reads consecutive chunks from the storage
func (p *simplePossumProvider) ReadConsecutiveChunks(prefix string) (io.ReadCloser, error) {
	// In a real implementation, this would use the possum library
	// For now, we'll just return an empty reader
	return io.NopCloser(strings.NewReader("")), nil
}

// Close closes the provider
func (p *simplePossumProvider) Close() error {
	return nil
}
