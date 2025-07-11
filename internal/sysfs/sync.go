//go:build !windows

package sysfs

import (
	"os"

	"github.com/yokecd/wazero/experimental/sys"
)

func fsync(f *os.File) sys.Errno {
	return sys.UnwrapOSError(f.Sync())
}
