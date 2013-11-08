// Copyright 2014 Tom Grennan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fixme_test

import (
	"bytes"
	"github.com/tgrennan/fixme"
	"testing"
)

func Test(t *testing.T) {
	var got bytes.Buffer
	fixme.SetWriter(&got)
	if fixme.Print("hello"); got.Len() != 0 {
		t.Error("pre-enable")
	}
	fixme.Enable()
	if fixme.Print("abc", "123"); got.String() !=
		"fixme_test.go:20: FIXME abc123\n" {
		t.Error("got:", got.String())
	}
	got.Reset()
	if fixme.Println("hello", "world"); got.String() !=
		"fixme_test.go:25: FIXME hello world\n" {
		t.Error("got:", got.String())
	}
	got.Reset()
	fixme.SetFlags(0)
	if fixme.Print("abc", "123"); got.String() !=
		"FIXME abc123\n" {
		t.Error("got:", got.String())
	}
	got.Reset()
	if fixme.Println("hello", "world"); got.String() !=
		"FIXME hello world\n" {
		t.Error("got:", got.String())
	}
}
