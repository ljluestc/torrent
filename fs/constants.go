package torrentfs
package torrentfs
package torrentfs
package torrentfs

import (
	"os"
)

const (
	// Default mode for directories
	defaultMode = os.FileMode(0755)
	
	// Default mode for files
	fileMode = os.FileMode(0644)
	
	// Size of a block for read operations
	readBlockSize = 32 * 1024
)
import (
	"os"
)

const (
	// DefaultMode is the file mode used for directories in the filesystem
	defaultMode = 0555 | os.ModeDir
	
	// DefaultFileMode is the file mode used for regular files
	defaultFileMode = 0444
	
	// DefaultReadaheadPieces is the default number of pieces to read ahead
	defaultReadaheadPieces = 5
	
	// DefaultUnmountTimeout is the timeout for unmounting a filesystem
	defaultUnmountTimeout = 5 // seconds
)
import (
	"os"
)

const (
	// DefaultMode is the file mode used for files in the filesystem
	defaultMode = 0555 | os.ModeDir
	
	// DefaultFileMode is the file mode used for regular files
	defaultFileMode = 0444
	
	// DefaultCacheCapacity is the default number of pieces to cache
	defaultCacheCapacity = 50
	
	// DefaultReadaheadPieces is the default number of pieces to read ahead
	defaultReadaheadPieces = 5
)
import (
	"os"
)

// Constants used across the torrentfs package
var (
	defaultMode = os.FileMode(0555)
)
