package storage

import (
	"github.com/anacrolix/torrent/metainfo"
)

// TorrentImpl is a concrete implementation of the TorrentImpl interface.
type TorrentImpl struct {
	Piece    func(p metainfo.Piece) PieceImpl
	Close    func() error
	Flush    func() error
	Capacity *func() (int64, bool)
}

// Piece calls the Piece function.
func (me TorrentImpl) Piece(p metainfo.Piece) PieceImpl {
	return me.Piece(p)
}

// PieceWithHash implements the interface with piece hash.
func (me TorrentImpl) PieceWithHash(p metainfo.Piece, pieceHash []byte) PieceImpl {
	// Just ignore hash if not supported
	return me.Piece(p)
}

// Close calls the Close function.
func (me TorrentImpl) Close() error {
	if me.Close == nil {
		return nil
	}
	return me.Close()
}
