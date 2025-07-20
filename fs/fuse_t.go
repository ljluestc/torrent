//go:build linux

package torrentfs

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/anacrolix/fuse"
	fusefs "github.com/anacrolix/fuse/fs"
)

// IsFuseTAvailable checks if fuse-t is available on the system
func IsFuseTAvailable() bool {
	// This can be replaced with a build tag (see fuse_t_build.go)
	return false
}

// FuseMounter defines an interface for mounting/unmounting a filesystem
type FuseMounter interface {
	Mount(fs fusefs.FS, mountpoint string) error
	Unmount(mountpoint string) error
}

// DefaultFuseMounter uses the standard anacrolix/fuse implementation
type DefaultFuseMounter struct{}

func (m DefaultFuseMounter) Mount(fs fusefs.FS, mountpoint string) error {
	err := os.MkdirAll(mountpoint, 0755)
	if err != nil {
		return fmt.Errorf("error creating mount point: %w", err)
	}

	conn, err := fuse.Mount(
		mountpoint,
		fuse.FSName("torrentfs"),
		fuse.Subtype("torrentfs"),
		fuse.LocalVolume(),
		fuse.VolumeName("TorrentFS"),
	)
	if err != nil {
		return err
	}

	go func() {
		err := fusefs.Serve(conn, fs)
		if err != nil {
			fmt.Printf("Error serving FUSE: %v\n", err)
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

func (m DefaultFuseMounter) Unmount(mountpoint string) error {
	return fuse.Unmount(mountpoint)
}

// GetFuseMounter returns the appropriate mounter for the platform
func GetFuseMounter() FuseMounter {
	if IsFuseTAvailable() {
		// Return the fuse-t mounter if available
		// This will be implemented in fuse_t_build.go
		return getFuseTMounter()
	}
	return DefaultFuseMounter{}
}

// mountTracker keeps track of mounted filesystems
var mountTracker struct {
	sync.Mutex
	mountpoints map[string]struct{}
}

func init() {
	mountTracker.mountpoints = make(map[string]struct{})
}

// MountFS mounts a filesystem using the appropriate mounter
func MountFS(fs fusefs.FS, mountpoint string) error {
	mounter := GetFuseMounter()
	err := mounter.Mount(fs, mountpoint)
	if err != nil {
		return err
	}

	// Track the mount for cleanup
	mountTracker.Lock()
	mountTracker.mountpoints[mountpoint] = struct{}{}
	mountTracker.Unlock()

	return nil
}

// UnmountFS unmounts a filesystem
func UnmountFS(mountpoint string) error {
	mounter := GetFuseMounter()
	err := mounter.Unmount(mountpoint)

	// Remove from tracked mounts
	mountTracker.Lock()
	delete(mountTracker.mountpoints, mountpoint)
	mountTracker.Unlock()

	return err
}

// CleanupAllMounts unmounts all tracked mounts
func CleanupAllMounts() {
	mountTracker.Lock()
	defer mountTracker.Unlock()

	for mp := range mountTracker.mountpoints {
		_ = GetFuseMounter().Unmount(mp)
		delete(mountTracker.mountpoints, mp)
	}
}

// Return a stub fuse-t mounter for non-fuset builds
func getFuseTMounter() FuseMounter {
	return DefaultFuseMounter{}
}
package torrentfs

import (
	"context"
	"errors"
	"os"
	"sync"
	"time"

	"github.com/anacrolix/fuse"
	fusefs "github.com/anacrolix/fuse/fs"
	"github.com/fuse-t/fuse-t/fuseops"
	fusetsys "github.com/fuse-t/fuse-t/fuseutil"
)

// FuseTAdapter is an adapter that connects the anacrolix/fuse API to fuse-t
type FuseTAdapter struct {
	fs            fusefs.FS
	inodes        map[fuseops.InodeID]fusefs.Node
	nextInodeID   fuseops.InodeID
	handles       map[fuseops.HandleID]fusefs.Handle
	nextHandleID  fuseops.HandleID
	rootNode      fusefs.Node
	rootAttrs     fuse.Attr
	rootDirHandle fusefs.Handle
	mu            sync.Mutex
}

// NewFuseTAdapter creates a new adapter that wraps an anacrolix/fuse FS
func NewFuseTAdapter(fs fusefs.FS) (*FuseTAdapter, error) {
	adapter := &FuseTAdapter{
		fs:           fs,
		inodes:       make(map[fuseops.InodeID]fusefs.Node),
		nextInodeID:  fuseops.InodeID(2), // Root is 1
		handles:      make(map[fuseops.HandleID]fusefs.Handle),
		nextHandleID: fuseops.HandleID(1),
	}

	// Get the root node
	var err error
	ctx := context.Background()
	adapter.rootNode, err = fs.Root()
	if err != nil {
		return nil, err
	}

	// Get root attributes
	if err := adapter.rootNode.Attr(ctx, &adapter.rootAttrs); err != nil {
		return nil, err
	}

	// Store root inode
	adapter.inodes[fuseops.InodeID(1)] = adapter.rootNode

	return adapter, nil
}

// Mount implements mounting the filesystem using fuse-t
func (a *FuseTAdapter) Mount(mountPoint string) error {
	server := fusetsys.NewFileSystemServer(a)
	return server.Mount(mountPoint)
}
//go:build !notorrentfs
// +build !notorrentfs

package torrentfs

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/anacrolix/fuse"
	fusefs "github.com/anacrolix/fuse/fs"
)

// FuseTAvailable indicates whether fuse-t support is available
var FuseTAvailable = false

// FuseMounter defines the interface for mounting a filesystem
type FuseMounter interface {
	Mount(mountPoint string) error
	Unmount(mountPoint string) error
}

// defaultFuseMounter uses the standard fuse library
type defaultFuseMounter struct{}

func (d defaultFuseMounter) Mount(mountPoint string) error {
	fs := New(nil) // This will be replaced in actual usage
	fuseCfg := fuse.Config{}
	conn, err := fuse.Mount(mountPoint, fusefs.New(fs), &fuseCfg)
	if err != nil {
		return fmt.Errorf("error mounting: %w", err)
	}
	go conn.Serve()
	return nil
}

func (d defaultFuseMounter) Unmount(mountPoint string) error {
	return fuse.Unmount(mountPoint)
}

// GetDefaultMounter returns the appropriate mounter for the platform
func GetDefaultMounter() FuseMounter {
	// When using fuse-t, this would be replaced with a fuse-t implementation
	return defaultFuseMounter{}
}

// IsFuseTAvailable checks if fuse-t is available on the system
func IsFuseTAvailable() bool {
	return FuseTAvailable
}
//go:build !notorrentfs && !fuset
// +build !notorrentfs,!fuset

package torrentfs

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/anacrolix/fuse"
	fusefs "github.com/anacrolix/fuse/fs"
)

// IsFuseTAvailable checks if fuse-t is available on the system
// When using the fuset build tag, this will be overridden
var IsFuseTAvailable = false

// defaultMode is the default file mode for directories
const defaultMode = 0755

// FuseMounter defines an interface for mounting/unmounting a filesystem
type FuseMounter interface {
	Mount(fs fusefs.FS, mountpoint string) error
	Unmount(mountpoint string) error
}

// DefaultFuseMounter uses the standard anacrolix/fuse implementation
type DefaultFuseMounter struct{}

func (m DefaultFuseMounter) Mount(fs fusefs.FS, mountpoint string) error {
	err := os.MkdirAll(mountpoint, 0755)
	if err != nil {
		return fmt.Errorf("error creating mount point: %w", err)
	}

	conn, err := fuse.Mount(
		mountpoint,
		fuse.FSName("torrentfs"),
		fuse.Subtype("torrentfs"),
		fuse.LocalVolume(),
		fuse.VolumeName("TorrentFS"),
	)
	if err != nil {
		return err
	}

	go func() {
		err := fusefs.Serve(conn, fs)
		if err != nil {
			fmt.Printf("Error serving FUSE: %v\n", err)
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

func (m DefaultFuseMounter) Unmount(mountpoint string) error {
	return fuse.Unmount(mountpoint)
}

// GetFuseMounter returns the appropriate mounter for the platform
func GetFuseMounter() FuseMounter {
	// Default implementation, will be overridden when using fuset build tag
	return DefaultFuseMounter{}
}

// mountTracker keeps track of mounted filesystems
var mountTracker struct {
	sync.Mutex
	mountpoints map[string]struct{}
}

func init() {
	mountTracker.mountpoints = make(map[string]struct{})
}

// MountFS mounts a filesystem using the appropriate mounter
func MountFS(fs fusefs.FS, mountpoint string) error {
	mounter := GetFuseMounter()
	err := mounter.Mount(fs, mountpoint)
	if err != nil {
		return err
	}

	// Track the mount for cleanup
	mountTracker.Lock()
	mountTracker.mountpoints[mountpoint] = struct{}{}
	mountTracker.Unlock()

	return nil
}

// UnmountFS unmounts a filesystem
func UnmountFS(mountpoint string) error {
	mounter := GetFuseMounter()
	err := mounter.Unmount(mountpoint)

	// Remove from tracked mounts
	mountTracker.Lock()
	delete(mountTracker.mountpoints, mountpoint)
	mountTracker.Unlock()

	return err
}

// CleanupAllMounts unmounts all tracked mounts
func CleanupAllMounts() {
	mountTracker.Lock()
	defer mountTracker.Unlock()

	for mp := range mountTracker.mountpoints {
		_ = GetFuseMounter().Unmount(mp)
		delete(mountTracker.mountpoints, mp)
	}
}

// IsFuseTAvailable checks if fuse-t is available on the system
func IsFuseTAvailable() bool {
	return FuseTAvailable
}
// MountFuse mounts a filesystem using either standard FUSE or fuse-t
func MountFuse(fs fusefs.FS, mountPoint string) error {
	// Create directories if they don't exist
	err := os.MkdirAll(mountPoint, 0755)
	if err != nil {
		return fmt.Errorf("error creating mount point: %w", err)
	}

	// Use the standard FUSE library
	cfg := fuse.Config{}
	conn, err := fuse.Mount(mountPoint, fusefs.New(fs), &cfg)
	if err != nil {
		return fmt.Errorf("error mounting: %w", err)
	}
	go conn.Serve()

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

// UnmountFuse unmounts a filesystem
func UnmountFuse(mountPoint string) error {
	return fuse.Unmount(mountPoint)
}

// fuseTState tracks the mount state for cleanup
var fuseTState struct {
	sync.Mutex
	mounts map[string]bool
}

func init() {
	fuseTState.mounts = make(map[string]bool)
}

// RegisterMount adds a mount point to the tracking state
func RegisterMount(mountPoint string) {
	fuseTState.Lock()
	defer fuseTState.Unlock()
	fuseTState.mounts[mountPoint] = true
}

// UnregisterMount removes a mount point from the tracking state
func UnregisterMount(mountPoint string) {
	fuseTState.Lock()
	defer fuseTState.Unlock()
	delete(fuseTState.mounts, mountPoint)
}

// CleanupAllMounts unmounts all tracked mount points
func CleanupAllMounts() {
	fuseTState.Lock()
	mountPoints := make([]string, 0, len(fuseTState.mounts))
	for mp := range fuseTState.mounts {
		mountPoints = append(mountPoints, mp)
	}
	fuseTState.Unlock()

	for _, mp := range mountPoints {
		_ = UnmountFuse(mp)
		UnregisterMount(mp)
	}
}
// StatFS implements the StatFS fuseops method
func (a *FuseTAdapter) StatFS(ctx context.Context, op *fuseops.StatFSOp) error {
	// Provide default values
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

// LookUpInode implements the LookUpInode fuseops method
func (a *FuseTAdapter) LookUpInode(ctx context.Context, op *fuseops.LookUpInodeOp) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Get parent node
	parentNode, ok := a.inodes[op.Parent]
	if !ok {
		return errors.New("parent node not found")
	}

	// Look up the child in the parent directory
	lookupFS, ok := parentNode.(fusefs.NodeStringLookuper)
	if !ok {
		return errors.New("parent node does not support lookup")
	}

	// Perform lookup
	childNode, err := lookupFS.Lookup(ctx, op.Name)
	if err != nil {
		return err
	}

	// Get child attributes
	var attr fuse.Attr
	if err := childNode.Attr(ctx, &attr); err != nil {
		return err
	}

	// Allocate a new inode ID for this child
	childID := a.nextInodeID
	a.nextInodeID++
	a.inodes[childID] = childNode

	// Set the response fields
	op.Entry.Child = childID
	op.Entry.Attributes.Mode = attr.Mode
	op.Entry.Attributes.Size = attr.Size
	op.Entry.Attributes.Atime = time.Unix(int64(attr.Atime), 0)
	op.Entry.Attributes.Mtime = time.Unix(int64(attr.Mtime), 0)
	op.Entry.Attributes.Ctime = time.Unix(int64(attr.Ctime), 0)
	op.Entry.Attributes.Uid = attr.Uid
	op.Entry.Attributes.Gid = attr.Gid

	return nil
}

// GetInodeAttributes implements the GetInodeAttributes fuseops method
func (a *FuseTAdapter) GetInodeAttributes(ctx context.Context, op *fuseops.GetInodeAttributesOp) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	node, ok := a.inodes[op.Inode]
	if !ok {
		return errors.New("inode not found")
	}

	var attr fuse.Attr
	if err := node.Attr(ctx, &attr); err != nil {
		return err
	}

	op.Attributes.Mode = attr.Mode
	op.Attributes.Size = attr.Size
	op.Attributes.Atime = time.Unix(int64(attr.Atime), 0)
	op.Attributes.Mtime = time.Unix(int64(attr.Mtime), 0)
	op.Attributes.Ctime = time.Unix(int64(attr.Ctime), 0)
	op.Attributes.Uid = attr.Uid
	op.Attributes.Gid = attr.Gid

	return nil
}

// OpenFile implements the OpenFile fuseops method
func (a *FuseTAdapter) OpenFile(ctx context.Context, op *fuseops.OpenFileOp) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	node, ok := a.inodes[op.Inode]
	if !ok {
		return errors.New("inode not found")
	}

	// Check if node implements the NodeOpener interface
	opener, ok := node.(fusefs.NodeOpener)
	if !ok {
		return errors.New("node does not support open")
	}

	// Create fuse open request
	req := &fuse.OpenRequest{
		Flags: fuse.OpenFlags(op.OpenFlags),
	}
	resp := &fuse.OpenResponse{}

	// Open the file
	handle, err := opener.Open(ctx, req, resp)
	if err != nil {
		return err
	}

	// Allocate a handle ID
	handleID := a.nextHandleID
	a.nextHandleID++
	a.handles[handleID] = handle
	op.Handle = handleID

	return nil
}

// ReadFile implements the ReadFile fuseops method
func (a *FuseTAdapter) ReadFile(ctx context.Context, op *fuseops.ReadFileOp) error {
	a.mu.Lock()
	handle, ok := a.handles[op.Handle]
	a.mu.Unlock()

	if !ok {
		return errors.New("handle not found")
	}

	// Check if handle implements the HandleReader interface
	reader, ok := handle.(fusefs.HandleReader)
	if !ok {
		return errors.New("handle does not support read")
	}

	// Create fuse read request
	req := &fuse.ReadRequest{
		Offset: op.Offset,
		Size:   op.Size,
	}
	resp := &fuse.ReadResponse{
		Data: make([]byte, op.Size),
	}

	// Read from the file
	if err := reader.Read(ctx, req, resp); err != nil {
		return err
	}

	// Copy data to the operation buffer
	copy(op.Data, resp.Data)
	op.BytesRead = len(resp.Data)

	return nil
}

// ReleaseFileHandle implements the ReleaseFileHandle fuseops method
func (a *FuseTAdapter) ReleaseFileHandle(ctx context.Context, op *fuseops.ReleaseFileHandleOp) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	handle, ok := a.handles[op.Handle]
	if !ok {
		return errors.New("handle not found")
	}

	// Check if handle implements the HandleReleaser interface
	if releaser, ok := handle.(fusefs.HandleReleaser); ok {
		if err := releaser.Release(ctx, &fuse.ReleaseRequest{}); err != nil {
			return err
		}
	}

	// Remove the handle
	delete(a.handles, op.Handle)
	return nil
}

// Mount a filesystem using fuse-t
func MountWithFuseT(fs fusefs.FS, mountpoint string) error {
	adapter, err := NewFuseTAdapter(fs)
	if err != nil {
		return err
	}
	return adapter.Mount(mountpoint)
}

// Unmount a filesystem mounted with fuse-t
func UnmountWithFuseT(mountpoint string) error {
	return fusetsys.Unmount(mountpoint)
}

// IsFuseTAvailable checks if fuse-t is available on the system
func IsFuseTAvailable() bool {
	// Check if we can import fuse-t and if it's actually usable
	return true // This is simplified - in reality you'd need to check if it works on the platform
}
