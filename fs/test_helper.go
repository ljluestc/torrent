package torrentfs
package torrentfs

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent"
)

var (
	skipTestUnmountWedged = false
)

// SkipUnmountWedgedTest marks the TestUnmountWedged test to be skipped
func SkipUnmountWedgedTest() {
	skipTestUnmountWedged = true
}

// SetupTestClient creates a torrent client for testing
func SetupTestClient(t *testing.T) *torrent.Client {
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = filepath.Join(os.TempDir(), "torrentfs-test")
	cfg.DisableTrackers = true
	cfg.NoDHT = true
	cfg.DisableTCP = true
	cfg.DisableUTP = true
	cfg.Logger = log.Default.WithLevel(log.Debug)
	
	client, err := torrent.NewClient(cfg)
	if err != nil {
		t.Fatalf("Error creating torrent client: %v", err)
	}
	
	return client
}

// MountTestFS mounts a filesystem for testing and returns the mount point and unmount function
func MountTestFS(t *testing.T, fs interface{}) (mountpoint string, unmount func() error, err error) {
	// Create a temporary directory for mounting
	mountpoint, err = os.MkdirTemp("", "torrentfs-test-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating temp dir: %w", err)
	}

	// Mount the filesystem
	err = MountFS(fs, mountpoint)
	if err != nil {
		os.RemoveAll(mountpoint)
		return "", nil, fmt.Errorf("error mounting filesystem: %w", err)
	}

	// Return a function to unmount and clean up
	unmount = func() error {
		err := UnmountFS(mountpoint)
		if err != nil {
			return err
		}
		return os.RemoveAll(mountpoint)
	}

	return mountpoint, unmount, nil
}

// WaitForMount waits for a filesystem to be mounted and ready
func WaitForMount(ctx context.Context, mountpoint string) error {
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	for {
		_, err := os.Stat(mountpoint)
		if err == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for mount: %w", ctx.Err())
		case <-time.After(100 * time.Millisecond):
		}
	}
}

// IsCIEnvironment checks if the code is running in a CI environment
func IsCIEnvironment() bool {
	return os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != ""
}

// SkipOnUnsupportedOS skips the test if running on an unsupported OS
func SkipOnUnsupportedOS(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping test on Windows - FUSE not well supported")
	}
	
	// Skip on macOS if not using fuse-t
	if runtime.GOOS == "darwin" && !IsFuseTAvailable() {
		t.Skip("Skipping test on macOS without fuse-t")
	}
}
import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/anacrolix/fuse"
	fusefs "github.com/anacrolix/fuse/fs"
	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent"
)

var (
	skipTestUnmountWedged = false
)

// SkipUnmountWedgedTest marks the TestUnmountWedged test to be skipped
func SkipUnmountWedgedTest() {
	skipTestUnmountWedged = true
}

// SetupTestClient creates a torrent client for testing
func SetupTestClient(t *testing.T) *torrent.Client {
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = filepath.Join(os.TempDir(), "torrentfs-test")
	cfg.DisableTrackers = true
	cfg.NoDHT = true
	cfg.DisableTCP = true
	cfg.DisableUTP = true
	cfg.Logger = log.Default.WithLevel(log.Debug)
	
	client, err := torrent.NewClient(cfg)
	if err != nil {
		t.Fatalf("Error creating torrent client: %v", err)
	}
	
	return client
}

// MountTestFS mounts a filesystem for testing and returns the mount point and unmount function
func MountTestFS(t *testing.T, fs fusefs.FS) (mountpoint string, unmount func() error, err error) {
	// Create a temporary directory for mounting
	mountpoint, err = os.MkdirTemp("", "torrentfs-test-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating temp dir: %w", err)
	}

	// Mount the filesystem
	err = MountFS(fs, mountpoint)
	if err != nil {
		os.RemoveAll(mountpoint)
		return "", nil, fmt.Errorf("error mounting filesystem: %w", err)
	}

	// Return a function to unmount and clean up
	unmount = func() error {
		err := UnmountFS(mountpoint)
		if err != nil {
			return err
		}
		return os.RemoveAll(mountpoint)
	}

	return mountpoint, unmount, nil
}

// WaitForMount waits for a filesystem to be mounted and ready
func WaitForMount(ctx context.Context, mountpoint string) error {
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	for {
		_, err := os.Stat(mountpoint)
		if err == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for mount: %w", ctx.Err())
		case <-time.After(100 * time.Millisecond):
		}
	}
}

// IsCIEnvironment checks if the code is running in a CI environment
func IsCIEnvironment() bool {
	return os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != ""
}

// SkipOnUnsupportedOS skips the test if running on an unsupported OS
func SkipOnUnsupportedOS(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping test on Windows - FUSE not well supported")
	}
	
	// Skip on macOS if not using fuse-t
	if runtime.GOOS == "darwin" && !IsFuseTAvailable() {
		t.Skip("Skipping test on macOS without fuse-t")
	}
}
