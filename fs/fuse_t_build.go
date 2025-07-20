//go:build fuset
// +build fuset

package torrentfs

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/anacrolix/fuse"
	fusefs "github.com/anacrolix/fuse/fs"
	"github.com/fuse-t/fuse-t/fuseops"
	fusetsys "github.com/fuse-t/fuse-t/fuseutil"
)

// FuseTMounter is the fuse-t implementation of the FuseMounter interface
type FuseTMounter struct{}

func (m FuseTMounter) Mount(fs fusefs.FS, mountpoint string) error {
	// Create the directory if it doesn't exist
	err := os.MkdirAll(mountpoint, 0755)
	if err != nil {
		return fmt.Errorf("error creating mount point: %w", err)
	}

	// Create an adapter to convert between anacrolix/fuse and fuse-t
	adapter := newFuseTAdapter(fs)
	
	// Mount the filesystem using fuse-t
	server := fusetsys.NewFileSystemServer(adapter)
	go func() {
		err := server.Mount(mountpoint)
		if err != nil {
			fmt.Printf("Error mounting with fuse-t: %v\n", err)
		}
	}()
	
	// Wait for the mount to be ready
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for {
		if _, err := os.Stat(mountpoint); err == nil {
			break
		}
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for mount: %w", ctx.Err())
		case <-time.After(100 * time.Millisecond):
		}
	}
	
	return nil
}

func (m FuseTMounter) Unmount(mountpoint string) error {
	return fusetsys.Unmount(mountpoint)
}

// IsFuseTAvailable returns true since this file is only built with the fuset tag
func IsFuseTAvailable() bool {
	return true
}
//go:build fuset
// +build fuset

package torrentfs

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/anacrolix/fuse"
	fusefs "github.com/anacrolix/fuse/fs"
	"github.com/fuse-t/fuse-t/fuseops"
	fusetsys "github.com/fuse-t/fuse-t/fuseutil"
)

func init() {
	// Enable fuse-t support when this file is built
	FuseTAvailable = true
}

// FuseTMounter is the fuse-t implementation of the FuseMounter interface
type FuseTMounter struct{}

func (m FuseTMounter) Mount(fs fusefs.FS, mountpoint string) error {
	// Create the directory if it doesn't exist
	err := os.MkdirAll(mountpoint, 0755)
	if err != nil {
		return fmt.Errorf("error creating mount point: %w", err)
	}

	// Create an adapter to convert between anacrolix/fuse and fuse-t
	adapter := newFuseTAdapter(fs)
	
	// Mount the filesystem using fuse-t
	server := fusetsys.NewFileSystemServer(adapter)
	go func() {
		err := server.Mount(mountpoint)
		if err != nil {
			fmt.Printf("Error mounting with fuse-t: %v\n", err)
		}
	}()
	
	// Wait for the mount to be ready
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for {
		if _, err := os.Stat(mountpoint); err == nil {
			break
		}
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for mount: %w", ctx.Err())
		case <-time.After(100 * time.Millisecond):
		}
	}
	
	return nil
}

func (m FuseTMounter) Unmount(mountpoint string) error {
	return fusetsys.Unmount(mountpoint)
}

// GetFuseMounter returns the fuse-t mounter when built with fuse-t support
func GetFuseMounter() FuseMounter {
	return FuseTMounter{}
}

// fuseTAdapter is a minimal adapter to use fuse-t with our existing code
type fuseTAdapter struct {
	fs fusefs.FS
}

// newFuseTAdapter creates a new adapter for fuse-t
func newFuseTAdapter(fs fusefs.FS) *fuseTAdapter {
	return &fuseTAdapter{fs: fs}
}

// Minimal implementation of the fuse-t filesystem interface
// In a complete implementation, this would need to implement all required methods
func (a *fuseTAdapter) StatFS(ctx context.Context, op *fuseops.StatFSOp) error {
	op.BlockSize = 4096
	op.Blocks = 1000000
	op.BlocksFree = 1000000
	op.BlocksAvailable = 1000000
	op.Files = 10000
	op.FilesFree = 10000
	op.IoSize = 4096
	op.MaxFileSize = 1 << 50
	return nil
}
// fuseTAdapter is a minimal adapter to use fuse-t with our existing code
type fuseTAdapter struct {
	fs fusefs.FS
}

// newFuseTAdapter creates a new adapter for fuse-t
func newFuseTAdapter(fs fusefs.FS) *fuseTAdapter {
	return &fuseTAdapter{fs: fs}
}

// Minimal implementation for demonstration purposes
// In a complete implementation, this would need to implement all the required methods
// from the fuse-t interface
func (a *fuseTAdapter) StatFS(ctx context.Context, op *fuseops.StatFSOp) error {
	op.BlockSize = 4096
	op.Blocks = 1000000
	op.BlocksFree = 1000000
	op.BlocksAvailable = 1000000
	op.Files = 10000
	op.FilesFree = 10000
	op.IoSize = 4096
	op.MaxFileSize = 1 << 50
	return nil
}

// GetFuseTMounter returns a fuse-t mounter
func getFuseTMounter() FuseMounter {
	return FuseTMounter{}
}
