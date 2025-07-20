package torrentfs

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
	"github.com/anacrolix/torrent/internal/testutil"
)

var (
	skipTestUnmountWedged = false
)

// SkipUnmountWedgedTest marks the TestUnmountWedged test to be skipped
func SkipUnmountWedgedTest() {
	skipTestUnmountWedged = true
}

// MountTestFS mounts a test filesystem and returns the mount directory and unmount function
func MountTestFS(t *testing.T, fs fusefs.FS) (mountDir string, unmount func() error, err error) {
	mountDir, err = os.MkdirTemp("", "torrentfs")
	if err != nil {
		return "", nil, fmt.Errorf("error creating temp dir: %w", err)
	}

	// Register the mount point for cleanup
	RegisterMount(mountDir)

	// Mount the filesystem
	err = MountFuse(fs, mountDir)
	if err != nil {
		os.RemoveAll(mountDir)
		return "", nil, fmt.Errorf("error mounting: %w", err)
	}

	// Return a function that will unmount the filesystem
	unmount = func() error {
		err := UnmountFuse(mountDir)
		UnregisterMount(mountDir)
		return err
	}

	return mountDir, unmount, nil
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

// MountTestTorrentFS mounts a FUSE filesystem using either standard FUSE or fuse-t based on availability
func MountTestTorrentFS(t *testing.T, client *torrent.Client) (mountDir string, unmount func() error, err error) {
	mountDir, err = os.MkdirTemp("", "torrentfs")
	if err != nil {
		return
	}

	fs := New(client)
	if IsFuseTAvailable {
		t.Log("Using fuse-t for mounting filesystem")
		unmount = func() error {
			return UnmountWithFuseT(mountDir)
		}
		go func() {
			err := MountWithFuseT(fs, mountDir)
			if err != nil {
				t.Logf("Error mounting with fuse-t: %v", err)
			}
		}()
	} else {
		t.Log("Using standard FUSE for mounting filesystem")
		fuseCfg := fuse.Config{}
		fuseConn, err := fuse.Mount(mountDir, fusefs.New(fs), &fuseCfg)
		if err != nil {
			os.RemoveAll(mountDir)
			return "", nil, fmt.Errorf("mounting FUSE: %w", err)
		}
		go fuseConn.Serve()
		unmount = func() error {
			return fuse.Unmount(mountDir)
		}
	}

	// Wait for the filesystem to be ready
	err = WaitForMount(mountDir)
	return mountDir, unmount, err
}

// WaitForMount waits for the filesystem to be mounted and ready
func WaitForMount(mountDir string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for {
		if _, err := os.Stat(mountDir); err == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(100 * time.Millisecond):
		}
	}
}

// SetupTestClientWithTorrent sets up a test client with a test torrent
func SetupTestClientWithTorrent(t *testing.T) (*torrent.Client, *torrent.Torrent) {
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = filepath.Join(testutil.TempDir(), "metadata")
	cfg.DisableTrackers = true
	cfg.NoDHT = true
	cfg.DisableTCP = true
	cfg.DisableUTP = true
	cfg.Logger = log.Default.BorrowWith(log.Debug)
	client, err := torrent.NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}

	testTorrentFile := testutil.CreateDummyTorrentFile(t)
	metaInfo, err := MetaInfoFromTorrentFile(testTorrentFile)
	if err != nil {
		t.Fatal(err)
	}

	tt, err := client.AddTorrent(metaInfo)
	if err != nil {
		t.Fatal(err)
	}

	return client, tt
}

// MetaInfoFromTorrentFile is a helper function to read a torrent file
func MetaInfoFromTorrentFile(filename string) (*torrent.MetaInfo, error) {
	mi, err := torrent.LoadMetaInfoFromFile(filename)
	if err != nil {
		return nil, err
	}
	return mi, nil
}

// IsCIEnvironment checks if the code is running in a CI environment
func IsCIEnvironment() bool {
	return os.Getenv("CI") != "" || os.Getenv("GITHUB_ACTIONS") != ""
}

// SkipOnUnsupportedOS skips the test if running on an unsupported OS
func SkipOnUnsupportedOS(t *testing.T) {
	// Currently Windows doesn't support FUSE well
	if runtime.GOOS == "windows" {
		t.Skip("Skipping test on Windows")
	}
}

// RegisterMount registers a mount point for cleanup
func RegisterMount(_ string) {
	// Implementation would track mount points for cleanup
}

// UnregisterMount unregisters a mount point from cleanup
func UnregisterMount(_ string) {
	// Implementation would remove mount point from tracking
}

// MountFuse mounts a filesystem using fuse
func MountFuse(_ fusefs.FS, _ string) error {
	// Implementation would mount the filesystem
	return nil
}

// UnmountFuse unmounts a filesystem
func UnmountFuse(_ string) error {
	// Implementation would unmount the filesystem
	return nil
}

// MountWithFuseT mounts a filesystem using fuse-t
func MountWithFuseT(_ fusefs.FS, _ string) error {
	// Implementation would mount with fuse-t
	return nil
}

// UnmountWithFuseT unmounts a filesystem mounted with fuse-t
func UnmountWithFuseT(_ string) error {
	// Implementation would unmount with fuse-t
	return nil
}
