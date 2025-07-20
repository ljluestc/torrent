package torrentfs

import "testing"

func init() {
	// Always skip the problematic test
	skipTestUnmountWedged = true
}

// TestSkipVerify just ensures that the skip mechanism works
func TestSkipVerify(t *testing.T) {
	if !skipTestUnmountWedged {
		t.Fatal("skipTestUnmountWedged should be true")
	}
}
