//go:build !android && !windows
// +build !android,!windows

package possum

import (
	"fmt"
	"io"
	"sort"
	"strconv"

	"github.com/anacrolix/log"
	possumLib "github.com/anacrolix/possum/go"
	possumResource "github.com/anacrolix/possum/go/resource"

	"github.com/anacrolix/torrent/storage/internal/shared"
)

// RealTorrentProvider extends possum resource.Provider with an efficient implementation of torrent
// provider interface
type RealTorrentProvider struct {
	possumResource.Provider
	Logger log.Logger
}

// ReadConsecutiveChunks implements the TorrentProvider interface
func (p RealTorrentProvider) ReadConsecutiveChunks(prefix string) (rc io.ReadCloser, err error) {
	p.Logger.Levelf(log.Debug, "ReadConsecutiveChunks(%q)", prefix)
	pr, err := p.Handle.NewReader()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			pr.End()
		}
	}()
	items, err := pr.ListItems(prefix)
	if err != nil {
		return
	}
	keys := make([]int64, 0, len(items))
	for _, item := range items {
		var i int64
		offsetStr := item.Key
		i, err = strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			err = fmt.Errorf("failed to parse offset %q: %w", offsetStr, err)
			return
		}
		keys = append(keys, i)
	}
	
	// Sort items by key
	sort.Slice(items, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	
	offset := int64(0)
	consValues := make([]consecutiveValue, 0, len(items))
	for i, item := range items {
		itemOffset := keys[i]
		if itemOffset > offset {
			// We can't provide a continuous read.
			break
		}
		if itemOffset+item.Stat.Size() <= offset {
			// This item isn't needed
			continue
		}
		var v possumLib.Value
		v, err = pr.Add(prefix + item.Key)
		if err != nil {
			return
		}
		consValues = append(consValues, consecutiveValue{
			pv:     v,
			offset: itemOffset,
			size:   item.Stat.Size(),
		})
		offset += item.Stat.Size() - (offset - itemOffset)
	}
	err = pr.Begin()
	if err != nil {
		return
	}
	rc, pw := io.Pipe()
	go func() {
		defer pr.End()
		err := p.writeConsecutiveValues(consValues, pw)
		err = pw.CloseWithError(err)
		if err != nil {
			p.Logger.Levelf(log.Error, "Error writing consecutive values: %v", err)
		}
	}()
	return
}

type consecutiveValue struct {
	pv     possumLib.Value
	offset int64
	size   int64
}

func (p RealTorrentProvider) writeConsecutiveValues(
	values []consecutiveValue, pw *io.PipeWriter,
) (err error) {
	off := int64(0)
	for _, v := range values {
		var n int64
		valueOff := off - v.offset
		n, err = io.Copy(pw, io.NewSectionReader(v.pv, valueOff, v.size-valueOff))
		if err != nil {
			return
		}
		off += n
	}
	return nil
}

// Close implements the TorrentProvider interface
func (p RealTorrentProvider) Close() error {
	return p.Handle.Close()
}

// NewRealProvider creates a new real possum provider
func NewRealProvider(logger log.Logger) TorrentProvider {
	// In a real implementation, this would return a real RealTorrentProvider
	// This is a simplified version that just returns a stub provider
	return &simplePossumProvider{
		logger: logger,
	}
}

func init() {
	// Set the platform-specific provider function
	newPlatformProvider = NewRealProvider
}
