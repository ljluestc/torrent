package storage

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/anacrolix/missinggo/v2"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// Adapts a ClientImpl to a Client.
func NewClient(impl ClientImpl) *Client {
	return &Client{
		impl: impl,
	}
}

// Wraps ClientImpl to provide a FilesStorage instance for a Torrent.
type Client struct {
	impl ClientImpl
}

// Opens a storage for the given torrent. If the torrent is unknown, a blank storage is returned
// that should allow pieceHashValid to work, and writes to be hashed to check against the expected
// results. The storage should not allow reads, to avoid serving bad data. 'info' is optional but
// recommended, allows implementations to check that the torrent file is as expected.
func (cl *Client) OpenTorrent(info *metainfo.Info, infoHash metainfo.Hash) (Torrent, error) {
	ctx := context.Background()
	it, err := cl.impl.OpenTorrent(ctx, info, infoHash)
	if err != nil {
		return nil, err
	}
	return &torrent{
		TorrentImpl: it,
	}, nil
}

// Deprecated: Use OpenTorrent instead.
func (cl *Client) OpenTorrentFile(torrent *metainfo.TorrentFile) (t Torrent, err error) {
	t, err = cl.OpenTorrent(torrent.Info(), torrent.HashInfoBytes())
	return
}

// Return a Client for the storage at the given path.
func NewFileWithCompletion(baseDir string) *Client {
	return NewClient(NewFileOpts(FileOpts{
		BaseDir:    baseDir,
		Completion: pieceCompletionForDir(baseDir),
	}))
}

func NewMMap(baseDir string) *Client {
	return NewClient(NewMMapWithCompletion(baseDir, pieceCompletionForDir(baseDir)))
}

// Returns the path to a file where piece completion pieces should be stored,
// without extension. A blank string means disable piece completion storage.
func pieceCompletionForDir(baseDir string) string {
	return filepath.Join(baseDir, ".torrent.bolt")
}

// Returns a fileStorageTorrent that uses the deprecated incomplete download
// directory.
func NewFileOpts(opts FileOpts) ClientImpl {
	if opts.NoCacheRead {
		panic("NoCacheRead may cause blocks to be written to storage with their default value. See godoc")
	}
	fc := &fileClientImpl{
		opts: opts,
	}
	missinggo.CopyExact(&fc.opts, &opts)
	return fc
}

// No need to cache reads when the storage has its own page cache.
func NewMMapWithCompletion(dir string, completion string) ClientImpl {
	return NewFileOpts(FileOpts{
		BaseDir:         dir,
		Completion:      completion,
		DisableRdFile:   true, // Required for MMap to receive ReadAt syscalls.
		SimpleBlobsPath: filepath.Join(dir, ".torrent.pieces"),
	})
}

// CloseClient closes the client, satisfying the storage.ClientCloser interface.
// TODO: Remove ClientCloser, it's superfluous.
func CloseClient(client *Client) error {
	if client.impl == nil {
		// TODO: log.Print("ClientImpl closed: nil implementation, was it closed before?")
	}
	impl := client.impl
	client.impl = nil
	// TODO: Is this if necessary?
	if impl == nil {
		return nil
	}
	if impl, ok := impl.(ClientImplCloser); ok {
		return impl.Close()
	}
	return nil
}

// I think these helpers are here to make writing implementation easier.

// type Pieces interface {
// 	GetPieceFromReaderAt(metainfo.Piece) PieceImpl
// 	ReadAt([]byte, int64)
// }

// type ResourceProvider interface {
// 	OpenTorrentResource(infoHash metainfo.Hash, res string) (_ io.ReadWriterAt, release func(), _ error)
// }

// Use a resource provider to get resources from infoHashes and paths.
type ResourcePiecesOpts struct {
	// Remove non-completion-token chunks from incomplete pieces. This allows space-constrained usage
	// models, and is particularly useful in Possum data storage because otherwise it has to use a
	// lot of space to track file positions.
	LeaveIncompleteChunks bool
}

// Adapts a Provider to a Client.
func NewResourcePieces(p interface {
	ReadResource(string) (io.ReadCloser, error)
	WriteResource(path string, f func(io.Writer) error) error
	GetLength(path string) (int64, error)
	Close() error
}) *resourcePiecesClientImpl {
	return NewResourcePiecesOpts(p, ResourcePiecesOpts{})
}

func NewResourcePiecesOpts(
	p interface {
		ReadResource(string) (io.ReadCloser, error)
		WriteResource(path string, f func(io.Writer) error) error
		GetLength(path string) (int64, error)
		Close() error
	},
	opts ResourcePiecesOpts,
) *resourcePiecesClientImpl {
	return &resourcePiecesClientImpl{
		p:    p,
		opts: opts,
	}
}

// Returns an io.ReaderAt if readable, otherwise returns an error. Can't use existing interfaces
// because ReadAt doesn't take io.Reader.
type Reader interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

var _ ClientImpl = &Client{}
