//go:build !stub
// +build !stub

package possum

import (
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// Piece implements PieceImpl
type Piece struct{}

// ReadAt implements PieceImpl
func (p *Piece) ReadAt(b []byte, off int64) (n int, err error) {
	return 0, nil
}

// WriteAt implements PieceImpl
func (p *Piece) WriteAt(b []byte, off int64) (n int, err error) {
	return len(b), nil
}

// MarkComplete implements PieceImpl
func (p *Piece) MarkComplete() error {
	return nil
}

// MarkNotComplete implements PieceImpl
func (p *Piece) MarkNotComplete() error {
	return nil
}

// Completion implements PieceImpl
func (p *Piece) Completion() shared.Completion {
	return shared.Completion{
		Complete: false,
		Ok:       true,
	}
}

// Ensure Piece implements shared.PieceImpl
var _ shared.PieceImpl = (*Piece)(nil)
