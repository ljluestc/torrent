//go:build android || windows
// +build android windows

package possum

import (
	"errors"
	"io"

	"github.com/anacrolix/log"
)

// StubTorrentProvider is a stub implementation of TorrentProvider
type StubTorrentProvider struct {
	logger log.Logger
}

// NewStubProvider creates a new stub provider
func NewStubProvider(logger log.Logger) *StubTorrentProvider {
	return &StubTorrentProvider{
		logger: logger,
	}
}

// ReadConsecutiveChunks returns an error as this is a stub implementation
func (p *StubTorrentProvider) ReadConsecutiveChunks(prefix string) (io.ReadCloser, error) {
	return nil, errors.New("possum is not supported on this platform")
}

// Close does nothing as this is a stub implementation
func (p *StubTorrentProvider) Close() error {
	return nil
}
