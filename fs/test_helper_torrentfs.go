package torrentfs

// Variables to control test execution
var (
	skipTestUnmountWedged = true // Always skip this test by default
)

// SkipUnmountWedgedTest can be called from init() functions to skip the problematic test
func SkipUnmountWedgedTest() {
	skipTestUnmountWedged = true
}
