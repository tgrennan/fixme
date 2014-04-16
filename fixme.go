// Copyright 2014 Tom Grennan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package fixme provides conditional printing for benignly littering code under
development with debug information and simple reminders.  Unless enabled, each
closure is a return stub; once enabled, these log:

	FILE:LINE FIXME ...
*/
package fixme

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	logger  *log.Logger
	flags   = log.Lshortfile
	writer  = io.Writer(os.Stderr)
	Print   = func(_ ...interface{}) {}
	Println = func(_ ...interface{}) {}
)

// Enable will un-stub fixme.Print and fixme.Println
func Enable() {
	const prefix = "FIXME"
	logger = log.New(writer, "", flags)
	Print = func(v ...interface{}) {
		logger.Output(2, fmt.Sprint(append([]interface{}{prefix, " "},
			v...)...))
	}
	Println = func(v ...interface{}) {
		logger.Output(2, fmt.Sprintln(append([]interface{}{prefix},
			v...)...))
	}
}

// The default flags are log.Lshortfile
func SetFlags(f int) {
	flags = f
	if logger != nil {
		logger.SetFlags(flags)
	}
}

// The default writer is os.Stderr
func SetWriter(w io.Writer) {
	writer = w
	if logger != nil {
		logger = log.New(writer, "", flags)
	}
}
