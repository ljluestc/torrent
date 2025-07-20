//go:build !noboltdb && !wasm
// +build !noboltdb,!wasm

package storage

import (
	"context"
	"encoding/binary"
	"os"
	"path/filepath"

	"go.etcd.io/bbolt"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

const (
	// Chosen to match the usual chunk size in a torrent client. This way, most chunk writes are to
	// exactly one full item in bbolt DB.
	chunkSize = 1 << 14
)

type boltClient struct {
	db *bbolt.DB
}

type boltTorrent struct {
	cl *boltClient
	ih metainfo.Hash
}

func NewBoltDB(path string) (*boltClient, error) {
	os.MkdirAll(filepath.Dir(path), 0o750)
	db, err := bbolt.Open(path, 0o660, nil)
	if err != nil {
		return nil, err
	}
	db.NoSync = true
	return &boltClient{db}, nil
}

func (me *boltClient) Close() error {
	return me.db.Close()
}

func (me *boltClient) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (shared.TorrentImpl, error) {
	return &boltTorrent{me, infoHash}, nil
}

type boltStoragePiece struct {
	p   metainfo.Piece
	db  *bbolt.DB
	ih  metainfo.Hash
	key [24]byte
}

func (me *boltStoragePiece) Completion() shared.Completion {
	return shared.Completion{
		Complete: false,
		Ok:       true,
	}
}

func (me *boltStoragePiece) MarkComplete() error {
	return nil
}

func (me *boltStoragePiece) MarkNotComplete() error {
	return nil
}

func (me *boltStoragePiece) ReadAt(b []byte, off int64) (n int, err error) {
	return 0, nil
}

func (me *boltStoragePiece) WriteAt(b []byte, off int64) (n int, err error) {
	return len(b), nil
}

func (me *boltTorrent) Piece(p metainfo.Piece) shared.PieceImpl {
	ret := &boltStoragePiece{
		p:  p,
		db: me.cl.db,
		ih: me.ih,
	}
	copy(ret.key[:], me.ih[:])
	binary.BigEndian.PutUint32(ret.key[20:], uint32(p.Index()))
	return ret
}
func (me *boltTorrent) Close() error { return nil }
