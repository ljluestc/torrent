// Package shared provides common interfaces used by both storage and storage/possum packages
// to avoid import cycles.
package shared

import (
	"io"

	"github.com/anacrolix/torrent/metainfo"
)

// Completion represents the completion status of a piece
type Completion struct {
	Complete bool
	Ok       bool
}

// TorrentImpl contains storage for a particular torrent.
type TorrentImpl interface {
	Piece(metainfo.Piece) PieceImpl
	Close() error
}

// PieceImpl is the storage for a particular piece.
type PieceImpl interface {
	io.ReaderAt
	io.WriterAt

	Completion() Completion
	MarkComplete() error
	MarkNotComplete() error
}

// ConsecutiveChunkReader interface allows reading consecutive chunks of data
type ConsecutiveChunkReader interface {
	ReadConsecutiveChunks(prefix string) (io.ReadCloser, error)
}
