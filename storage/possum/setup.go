// Package possum provides an implementation of torrent storage
// using the possum library.
//
// The package has two implementations:
// 1. A real implementation when built without the 'stub' build tag
// 2. A stub implementation when built with the 'stub' build tag
//
// For testing without the actual possum library, use:
//
//	go test -tags stub ./...
package possum

// Import this package with:
//   import "github.com/anacrolix/torrent/storage/possum"
