package storage

import (
	"context"
	"errors"
	"math/rand"
	"strings"

	"github.com/anacrolix/torrent/internal/testutil"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

type badStorage struct{}

var _ ClientImpl = badStorage{}

func (bs badStorage) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (shared.TorrentImpl, error) {
	return &badTorrent{}, nil
}

type badTorrent struct{}

func (bt *badTorrent) Piece(p metainfo.Piece) shared.PieceImpl {
	return badStoragePiece{p}
}

func (bt *badTorrent) Close() error {
	return nil
}

type badStoragePiece struct {
	p metainfo.Piece
}

var _ shared.PieceImpl = badStoragePiece{}

func (p badStoragePiece) WriteAt(b []byte, off int64) (int, error) {
	return 0, nil
}

func (p badStoragePiece) Completion() shared.Completion {
	return shared.Completion{Complete: true, Ok: true}
}

func (p badStoragePiece) MarkComplete() error {
	return errors.New("psyyyyyyyche")
}

func (p badStoragePiece) MarkNotComplete() error {
	return errors.New("psyyyyyyyche")
}

func (p badStoragePiece) randomlyTruncatedDataString() string {
	return testutil.GreetingFileContents[:rand.Intn(14)]
}

func (p badStoragePiece) ReadAt(b []byte, off int64) (n int, err error) {
	r := strings.NewReader(p.randomlyTruncatedDataString())
	return r.ReadAt(b, off+p.p.Offset())
}
