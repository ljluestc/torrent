package stub

import (
	"errors"

	"github.com/anacrolix/torrent/storage/internal/shared"
)

// Piece is a stub implementation for PieceImpl
type Piece struct{}

// ReadAt returns an error for stub implementation
func (p *Piece) ReadAt(b []byte, off int64) (n int, err error) {
	return 0, errors.New("stub implementation - not supported")
}

// WriteAt returns an error for stub implementation
func (p *Piece) WriteAt(b []byte, off int64) (n int, err error) {
	return 0, errors.New("stub implementation - not supported")
}

// MarkComplete returns an error for stub implementation
func (p *Piece) MarkComplete() error {
	return errors.New("stub implementation - not supported")
}

// MarkNotComplete returns an error for stub implementation
func (p *Piece) MarkNotComplete() error {
	return errors.New("stub implementation - not supported")
}

// Completion returns a default completion status
func (p *Piece) Completion() shared.Completion {
	return shared.Completion{
		Complete: false,
		Ok:       false,
	}
}
