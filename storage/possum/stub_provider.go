//go:build stub
// +build stub

package possumTorrentStorage

import (
	"io"

	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage"
)

// Provider is a stub implementation for the possum storage provider
type Provider struct {
	Provider interface{} // Just to match the real struct
	Logger   log.Logger
}

// NewClient creates a new client implementation
func (p Provider) NewClient() (storage.ClientImplCloser, error) {
	return &stubClient{
		logger: p.Logger,
	}, nil
}

// stubClient is a no-op client for testing
type stubClient struct {
	logger log.Logger
}

// Close implements storage.ClientImplCloser
func (c *stubClient) Close() error {
	return nil
}

// OpenTorrent implements storage.ClientImpl
func (c *stubClient) OpenTorrent(info *storage.TorrentInfo) (storage.TorrentImpl, error) {
	return &stubTorrent{
		info:   info,
		logger: c.logger,
	}, nil
}

// stubTorrent implements a stub storage.TorrentImpl
type stubTorrent struct {
	info   *storage.TorrentInfo
	logger log.Logger
}

// Piece implements storage.TorrentImpl
func (t *stubTorrent) Piece(p metainfo.Piece) storage.PieceImpl {
	return &stubPiece{}
}

// Close implements storage.TorrentImpl
func (t *stubTorrent) Close() error {
	return nil
}

// stubPiece implements a stub storage.PieceImpl
type stubPiece struct{}

// ReadAt implements storage.PieceImpl
func (p *stubPiece) ReadAt(b []byte, off int64) (n int, err error) {
	return 0, io.EOF
}

// WriteAt implements storage.PieceImpl
func (p *stubPiece) WriteAt(b []byte, off int64) (n int, err error) {
	return len(b), nil
}

// MarkComplete implements storage.PieceImpl
func (p *stubPiece) MarkComplete() error {
	return nil
}

// MarkNotComplete implements storage.PieceImpl
func (p *stubPiece) MarkNotComplete() error {
	return nil
}

// Completion implements storage.PieceImpl
func (p *stubPiece) Completion() storage.Completion {
	return storage.Completion{
		Complete: false,
		Ok:       true,
	}
}

// GetIsComplete implements the interface
func (p *stubPiece) GetIsComplete() bool {
	return false
}

var (
	_ storage.ClientImplCloser = &stubClient{}
	_ storage.TorrentImpl      = &stubTorrent{}
	_ storage.PieceImpl        = &stubPiece{}
)
