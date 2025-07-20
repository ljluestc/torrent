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
package possum

import (
	"context"
	"testing"

	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPossumClient(t *testing.T) {
	// Create a new client with a stub provider
	provider := NewProvider(log.Default)
	client := NewClient(provider)
	defer client.Close()

	// Test opening a torrent
	ctx := context.Background()
	torrentImpl, err := client.OpenTorrent(ctx, nil, metainfo.NewHashFromHex("0123456789abcdef0123456789abcdef01234567"))
	require.NoError(t, err)
	defer torrentImpl.Close()

	// Test getting a piece
	piece := torrentImpl.Piece(metainfo.Piece{})
	
	// Test piece completion
	completion := piece.Completion()
	assert.False(t, completion.Complete)
	assert.True(t, completion.Ok)

	// Test marking as complete
	err = piece.MarkComplete()
	assert.NoError(t, err)

	// Test reading from the piece
	buf := make([]byte, 10)
	n, err := piece.ReadAt(buf, 0)
	assert.Equal(t, 10, n)
	assert.NoError(t, err)

	// Test writing to the piece
	n, err = piece.WriteAt([]byte("test data"), 0)
	assert.Equal(t, 9, n)
	assert.NoError(t, err)
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
