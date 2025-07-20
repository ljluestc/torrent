package storage

import (
	"fmt"
	"io"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// TODO: Try and make this more readable.
type Torrent interface {
	// All Readers must be closed before closing the Torrent.
	Piece(metainfo.Piece) Piece
	Close() error
}

// torrent adapts a TorrentImpl to a Torrent.
type torrent struct {
	shared.TorrentImpl
}

// Ensure torrent implements Torrent interface
var _ Torrent = &torrent{}

func (t *torrent) Piece(p metainfo.Piece) Piece {
	return pieceImpl{
		t.TorrentImpl.Piece(p),
	}
}

func (t *torrent) Close() error {
	return t.TorrentImpl.Close()
}

// Implementations are not forced to support all methods on all piece types.
// Calling unsupported methods will panic.
type Piece interface {
	// These interfaces are not consolidated to make it easier for types to implement
	// only what they support.
	io.ReaderAt
	io.WriterAt
	PieceStatus() PieceStatus
	// Completion returns whether the piece is complete and the data is accessible
	// (probably by using ReadAt) in case of need.
	Completion() Completion
	// MarkComplete marks piece as complete, allowing it to be used as consistent
	// data source. Must be called after all checks are successful (hash, etc.).
	MarkComplete() error
	// MarkNotComplete returns an error if the piece status is now unknown.
	MarkNotComplete() error
}

// Storage for a piece. Implementations are not expected to be thread-safe.
type pieceImpl struct {
	PieceImpl shared.PieceImpl
}

var _ Piece = pieceImpl{}

func (s pieceImpl) MarkComplete() error {
	return s.PieceImpl.MarkComplete()
}

func (s pieceImpl) Completion() Completion {
	return s.PieceImpl.Completion()
}

func (s pieceImpl) MarkNotComplete() error {
	return s.PieceImpl.MarkNotComplete()
}

func (s pieceImpl) ReadAt(b []byte, off int64) (n int, err error) {
	return s.PieceImpl.ReadAt(b, off)
}

func (s pieceImpl) WriteAt(b []byte, off int64) (n int, err error) {
	return s.PieceImpl.WriteAt(b, off)
}

func (s pieceImpl) PieceStatus() PieceStatus {
	c := s.Completion()
	return PieceStatus{
		Complete: c.Complete,
		Err:      fmt.Sprintf("%v", c),
	}
}

// The PieceStatus gives a unified status of a piece.
type PieceStatus struct {
	Complete bool
	Err      string
}
