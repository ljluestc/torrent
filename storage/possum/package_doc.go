// Package possum provides a storage backend for torrents using possum.
//
// This package has two implementations:
// 1. The real implementation, used by default
// 2. A stub implementation, used when the "stub" build tag is specified
//
// To use the stub implementation, build or test with:
//
//	go build -tags stub
//	go test -tags stub
//
// All files in this package MUST use the same package name "possum" to prevent
// build errors related to multiple packages in the same directory.
package possum
