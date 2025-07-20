//go:build stub
// +build stub

package possumTorrentStorage

import (
	"testing"

	"github.com/anacrolix/log"
)

// TestProvider just ensures the storage provider can be instantiated
func TestProvider(t *testing.T) {
	provider := Provider{
		Logger: log.Default,
	}

	client, err := provider.NewClient()
	if err != nil {
		t.Fatalf("Error creating client: %v", err)
	}

	// Just make sure we can call Close without errors
	err = client.Close()
	if err != nil {
		t.Fatalf("Error closing client: %v", err)
	}
}
