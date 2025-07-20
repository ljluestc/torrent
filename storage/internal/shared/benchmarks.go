package shared

import (
	"context"
	"io"

	"github.com/anacrolix/torrent/metainfo"
)

// BenchmarkConsts contains constants for benchmarking
const (
	DefaultPieceSize = 1 << 14
	DefaultNumPieces = 1 << 6
)

// These interfaces allow testing without import cycles

// ClientImplBenchmark is a minimal interface for testing/benchmarking storage implementations
type ClientImplBenchmark interface {
	OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (TorrentImpl, error)
	Close() error
}

// ResourceProvider is an interface for providers that can read/write resources
type ResourceProvider interface {
	ReadResource(string) (io.ReadCloser, error)
	WriteResource(path string, f func(io.Writer) error) error
	GetLength(path string) (int64, error)
	Close() error
}

// Benchmark functions can be implemented here when needed
