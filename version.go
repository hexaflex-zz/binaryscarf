// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"runtime"
)

// Application name and version constants.
const (
	AppName         = "binaryscarf"
	AppVersionMajor = 0
	AppVersionMinor = 3
)

// Version returns the application version as a string.
func Version() string {
	return fmt.Sprintf("%s %d.%d (Go runtime %s).",
		AppName, AppVersionMajor, AppVersionMinor, runtime.Version())
}
