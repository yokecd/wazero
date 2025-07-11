//go:build !linux && !darwin && !windows

package sysfs

import (
	"github.com/yokecd/wazero/experimental/sys"
	"github.com/yokecd/wazero/internal/fsapi"
)

// poll implements `Poll` as documented on fsapi.File via a file descriptor.
func poll(uintptr, fsapi.Pflag, int32) (bool, sys.Errno) {
	return false, sys.ENOSYS
}
