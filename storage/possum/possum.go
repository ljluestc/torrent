package possum

import (
	"context"
	"io"
	"sync"

	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// Storage implements a storage backend for torrents using possum
type Storage struct {
	provider TorrentProvider
	logger   log.Logger
}

// New creates a new possum storage instance
func New(provider TorrentProvider) *Storage {
	return &Storage{
		provider: provider,
		logger:   log.Default,
	}
}

// WithLogger returns a new Storage with the given logger
func (s *Storage) WithLogger(logger log.Logger) *Storage {
	s.logger = logger
	return s
}

// OpenTorrent returns a new torrent instance
func (s *Storage) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (shared.TorrentImpl, error) {
	return &torrent{
		storage:  s,
		infoHash: infoHash,
		info:     info,
		prefix:   infoHash.HexString() + "/",
	}, nil
}

// Close closes the storage
func (s *Storage) Close() error {
	return s.provider.Close()
}

// Torrent implements shared.TorrentImpl
type torrent struct {
	storage  *Storage
	infoHash metainfo.Hash
	info     *metainfo.Info
	prefix   string
}

// Piece returns a piece implementation
func (t *torrent) Piece(p metainfo.Piece) shared.PieceImpl {
	return &piece{
		t:      t,
		piece:  p,
		prefix: t.prefix + "pieces/" + pieceKey(p.Index()),
	}
}

// Close closes the torrent
func (t *torrent) Close() error {
	return nil
}

// pieceKey returns a string key for a piece index
func pieceKey(index int) string {
	return string(index)
}

// Piece implements shared.PieceImpl
type piece struct {
	t      *torrent
	piece  metainfo.Piece
	prefix string
	mu     sync.Mutex
}

// WriteAt writes data to the piece
func (p *piece) WriteAt(b []byte, off int64) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// This is a simplified implementation
	// A real implementation would use the provider to write to storage
	// For now, pretend to write data
	return len(b), nil
}

// ReadAt reads data from the piece
func (p *piece) ReadAt(b []byte, off int64) (int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Try to read consecutive chunks
	// This is a simplified implementation
	rc, err := p.t.storage.provider.ReadConsecutiveChunks(p.prefix)
	if err != nil {
		// Fall back to filling with zeros
		for i := range b {
			b[i] = 0
		}
		return len(b), nil
	}
	defer rc.Close()

	// Seek to the offset
	if off > 0 {
		_, err = io.CopyN(io.Discard, rc, off)
		if err != nil {
			return 0, err
		}
	}

	// Read data
	return io.ReadFull(rc, b)
}

// Completion returns the completion status of the piece
func (p *piece) Completion() shared.Completion {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Check if the piece is complete
	// This is a simplified implementation
	return shared.Completion{
		Complete: false,
		Ok:       true,
	}
}

// MarkComplete marks the piece as complete
func (p *piece) MarkComplete() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Mark the piece as complete
	// This is a simplified implementation
	return nil
}

// MarkNotComplete marks the piece as not complete
func (p *piece) MarkNotComplete() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Mark the piece as not complete
	// This is a simplified implementation
	return nil
}

// PossumStorage is a mock storage implementation that pretends to store data
// but actually does nothing useful with it, like a possum playing dead.
type PossumStorage struct {
	mu sync.Mutex
}

// OpenTorrent returns a new torrent instance
func (ps *PossumStorage) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (shared.TorrentImpl, error) {
	return &possumTorrent{
		infoHash: infoHash,
	}, nil
}

// Close closes the storage
func (ps *PossumStorage) Close() error {
	return nil
}

// Torrent implementation for possum storage
type possumTorrent struct {
	infoHash metainfo.Hash
}

// Piece returns a piece implementation
func (pt *possumTorrent) Piece(p metainfo.Piece) shared.PieceImpl {
	return &possumPiece{
		piece: p,
	}
}

// Close closes the torrent
func (pt *possumTorrent) Close() error {
	return nil
}

// Piece implementation for possum storage
type possumPiece struct {
	piece     metainfo.Piece
	completed bool
	mu        sync.Mutex
}

// WriteAt pretends to write data
func (pp *possumPiece) WriteAt(b []byte, off int64) (int, error) {
	// Pretend to write all the data
	return len(b), nil
}

// ReadAt pretends to read data
func (pp *possumPiece) ReadAt(b []byte, off int64) (int, error) {
	// Fill the buffer with zeros
	for i := range b {
		b[i] = 0
	}
	return len(b), nil
}

// Completion returns the completion status
func (pp *possumPiece) Completion() shared.Completion {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	return shared.Completion{
		Complete: pp.completed,
		Ok:       true,
	}
}

// MarkComplete marks the piece as complete
func (pp *possumPiece) MarkComplete() error {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	pp.completed = true
	return nil
}

// MarkNotComplete marks the piece as not complete
func (pp *possumPiece) MarkNotComplete() error {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	pp.completed = false
	return nil
}
