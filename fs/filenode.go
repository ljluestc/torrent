package torrentfs

import (
	"context"

	"github.com/anacrolix/fuse"
	fusefs "github.com/anacrolix/fuse/fs"

	"github.com/anacrolix/torrent"
)

type fileNode struct {
	fusefs.Node
	tf *torrent.File
}

var _ fusefs.NodeOpener = (*fileNode)(nil)

func (fn *fileNode) Attr(ctx context.Context, attr *fuse.Attr) error {
	attr.Size = uint64(fn.tf.Length())
	attr.Mode = defaultMode
	return nil
}

func (fn *fileNode) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fusefs.Handle, error) {
	return &fileHandle{fn: fn, tf: fn.tf}, nil
}

type fileHandle struct {
	fn *fileNode
	tf *torrent.File
}

var _ fusefs.Handle = (*fileHandle)(nil)
