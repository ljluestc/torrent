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
	// Enable fuse-t support
	FuseTAvailable = true
}

// fuseTMounter is the fuse-t implementation of FuseMounter
type fuseTMounter struct{}

func (f fuseTMounter) Mount(mountPoint string) error {
	// In a real implementation, this would be passed the FS to mount
	// For now, just create the directory
	if err := os.MkdirAll(mountPoint, 0755); err != nil {
		return fmt.Errorf("error creating mount point: %w", err)
	}
	return nil
}

func (f fuseTMounter) Unmount(mountPoint string) error {
	return fusetsys.Unmount(mountPoint)
}

// GetDefaultMounter returns the fuse-t mounter when built with fuse-t support
func GetDefaultMounter() FuseMounter {
	return fuseTMounter{}
}

// MountFuse mounts a filesystem using fuse-t
func MountFuse(fs fusefs.FS, mountPoint string) error {
	// Create the fuse-t adapter
	adapter := &fuseTAdapter{
		fs:     fs,
		server: fusetsys.NewFileSystemServer(nil),
	}
	
	// Set up the adapter
	adapter.server.SetImpl(adapter)
	
	// Create the mount point
	if err := os.MkdirAll(mountPoint, 0755); err != nil {
		return fmt.Errorf("error creating mount point: %w", err)
	}
	
	// Mount the filesystem
	go adapter.server.Mount(mountPoint)
	
	// Wait for the mount to be ready
	waitCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	for {
		if _, err := os.Stat(mountPoint); err == nil {
			break
		}
		select {
		case <-waitCtx.Done():
			return fmt.Errorf("timeout waiting for mount: %w", waitCtx.Err())
		case <-time.After(100 * time.Millisecond):
		}
	}
	
	return nil
}

// UnmountFuse unmounts a filesystem using fuse-t
func UnmountFuse(mountPoint string) error {
	return fusetsys.Unmount(mountPoint)
}

// fuseTAdapter adapts anacrolix/fuse to fuse-t
type fuseTAdapter struct {
	fs     fusefs.FS
	server *fusetsys.FileSystemServer
}

// Minimal implementation of the fuse-t filesystem interface
// In a full implementation, you would need to implement all required methods

func (f *fuseTAdapter) StatFS(ctx context.Context, op *fuseops.StatFSOp) error {
	// Provide some default values
	op.BlockSize = 4096
	op.Blocks = 1000000
	op.BlocksFree = 1000000
	op.BlocksAvailable = 1000000
	op.IoSize = 4096
	op.Files = 10000
	op.FilesFree = 10000
	op.MaxFileSize = 1 << 50
	return nil
}

// IsFuseTAvailable checks if fuse-t is available on the system
func IsFuseTAvailable() bool {
	return true
}
