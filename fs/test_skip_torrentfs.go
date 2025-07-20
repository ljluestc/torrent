//go:build !notorrentfs && !fuset
// +build !notorrentfs,!fuset

package torrentfs

import (
	"os"
	"runtime"
)

// Skip torrentfs tests when fuse-t is not available
func init() {
	// Skip tests on CI if not using fuse-t build tag and not on platforms where regular FUSE works
	if (os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != "") &&
		(runtime.GOOS == "windows" || runtime.GOOS == "darwin") {
		SkipUnmountWedgedTest()
	}
}
