//go:build !possum
// +build !possum

package storage

import (
	"github.com/anacrolix/log"
	possumTorrentStorage "github.com/anacrolix/torrent/storage/possum"
)

// Return a fake possum provider so that this will compile
func NewPossum(log log.Logger) possumTorrentStorage.TorrentProvider {
	panic("not implemented")
}
