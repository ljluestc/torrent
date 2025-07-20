//go:build !stub
// +build !stub

package possum

import (
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// Torrent represents the torrent implementation
type Torrent struct {
	Logger      log.Logger
	TorrentImpl shared.TorrentImpl
	PieceReader func(metainfo.Piece) shared.PieceImpl
}

// Piece returns an implementation of a piece
func (t *Torrent) Piece(p metainfo.Piece) shared.PieceImpl {
	if t.PieceReader != nil {
		return t.PieceReader(p)
	}
	if t.TorrentImpl != nil {
		if impl, ok := t.TorrentImpl.(interface {
			Piece(metainfo.Piece) shared.PieceImpl
		}); ok {
			return impl.Piece(p)
		}
		return t.TorrentImpl.Piece(p)
	}
	return &Piece{}
}

// Close implements TorrentImpl
func (t *Torrent) Close() error {
	if t.TorrentImpl != nil {
		if closer, ok := t.TorrentImpl.(interface{ Close() error }); ok {
			return closer.Close()
		}
		return t.TorrentImpl.Close()
	}
	return nil
}

// Ensure Torrent implements shared.TorrentImpl
var _ shared.TorrentImpl = (*Torrent)(nil)
