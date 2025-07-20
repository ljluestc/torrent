package storage

import (
	"github.com/anacrolix/torrent/storage/internal/shared"
)

// Completion states the state of a piece.
type Completion = shared.Completion

// Returns true if complete, or not ok.
func (me Completion) IncompleteOk() bool {
	return !me.Complete && me.Ok
}
