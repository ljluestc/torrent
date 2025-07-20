//go:build stub
// +build stub

package torrentfs

// Skip TestUnmountWedged when using stub builds
func init() {
	SkipUnmountWedgedTest()
}
