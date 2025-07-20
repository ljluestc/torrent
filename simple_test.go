package torrent

import (
	"testing"
)

// TestSimple verifies that basic testing works
func TestSimple(t *testing.T) {
	// This test should always pass
	if 1+1 != 2 {
		t.Fatal("Basic math failed")
	}
}

// TestSimpleClientCreate verifies basic client creation
func TestSimpleClientCreate(t *testing.T) {
	config := NewDefaultClientConfig()
	client, err := NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
}
