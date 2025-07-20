package possum

import (
	"io"

	"github.com/anacrolix/log"
)

// TorrentProvider defines an interface for a provider that can read consecutive chunks
type TorrentProvider interface {
	ReadConsecutiveChunks(prefix string) (io.ReadCloser, error)
	Close() error
}

// NewProvider creates a new TorrentProvider based on the platform
func NewProvider(logger log.Logger) TorrentProvider {
	// The implementation is provided by platform-specific files
	return newPlatformProvider(logger)
}

// Platform-specific function that will be provided by the appropriate build file
var newPlatformProvider = func(logger log.Logger) TorrentProvider {
	// This will be replaced by the platform-specific implementation
	panic("no platform provider implementation available")
}
