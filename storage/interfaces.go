package storage

import (
	"context"
	"io"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// ClientImpl interface allows plugging in different storage implementations.
type ClientImpl interface {
	OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (shared.TorrentImpl, error)
}

// ClientImplCloser is an optional interface for ClientImpl that allows closing resources.
type ClientImplCloser interface {
	ClientImpl
	Close() error
}

// Type aliases from internal/shared for backward compatibility
type Completion = shared.Completion
type TorrentImpl = shared.TorrentImpl
type PieceImpl = shared.PieceImpl

// ConsecutiveChunkReader interface allows reading consecutive chunks of data
type ConsecutiveChunkReader interface {
	ReadConsecutiveChunks(prefix string) (io.ReadCloser, error)
}

// FileOpts contains options for file-based storage
type FileOpts struct {
	// Root directory for storing downloads.
	BaseDir string
	// The path for the piece completion database. Can be blank for no persistence. If
	// relative, it's relative to BaseDir.
	Completion string
	// Uses fewer resources, in particular by using mmap where possible.
	DisableRdFile bool
	// Disable preallocation of file space, which results in fragmentation and
	// additional disk activity while writing.
	NoPreallocate bool
	// Path to use for storing pieces that aren't a complete file.
	SimpleBlobsPath string
	// Only read and don't try write or create.
	ReadOnly bool
	// If true, don't use OS cache
	NoCacheRead bool
	// Piece completion database
	PieceCompletion PieceCompletion
}

// PieceCompletion represents piece completion database
type PieceCompletion interface {
	Get(metainfo.PieceKey) (bool, error)
	Set(metainfo.PieceKey, bool) error
	Close() error
}

// Torrent represents a storage implementation for a torrent
type Torrent interface {
	Piece(p metainfo.Piece) PieceImpl
	Close() error
}
