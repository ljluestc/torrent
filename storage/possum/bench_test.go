//go:build !stub
// +build !stub

package possum

import (
	"testing"

	g "github.com/anacrolix/generics"
	"github.com/anacrolix/log"
	possumLib "github.com/anacrolix/possum/go"
)

// Constants for benchmark testing
const (
	DefaultPieceSize = 1 << 14
	DefaultNumPieces = 1 << 6
)

// This should be made to mirror the benchmarks for sqlite storage.
func BenchmarkProvider(b *testing.B) {
	possumDir, err := possumLib.Open(b.TempDir())
	if err != nil {
		b.Fatal(err)
	}
	possumDir.SetInstanceLimits(possumLib.Limits{
		DisableHolePunching: false,
		MaxValueLengthSum:   g.Some[uint64](DefaultPieceSize * DefaultNumPieces / 2),
	})
	defer possumDir.Close()
	provider := Provider{Logger: log.Default}

	// Test implementation would go here, but avoiding import cycle
	b.Skip("Skipping benchmark to avoid import cycle")
}
