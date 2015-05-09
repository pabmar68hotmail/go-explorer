// Copyright 2015 Gary Burd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"strings"
	"testing"
)

var formatTests = []struct {
	in      string
	out     string
	fmtonly bool
}{
	{
		in:  "package main",
		out: "OK",
	},
	{
		// missing line between package and var
		in:  "package main\nvar i int",
		out: "REPL 2 1\n",
	},
	{
		// extra line at end
		in:  "package main\n",
		out: "REPL 2 2",
	},
	{
		// extra lines at end
		in:  "package main\n\n",
		out: "REPL 2 3",
	},
	{
		// extra line at start
		in:  "\npackage main",
		out: "REPL 1 1",
	},
	{
		// extra lines at start
		in:  "\n\n\npackage main",
		out: "REPL 1 3",
	},
	{
		// extra space
		in:  "package  main",
		out: "REPL 1 1\npackage main",
	},
	{
		// modify all lines
		in:  "package  main\n ",
		out: "REPL 1 2\npackage main",
	},
	{
		// modify all lines
		in:  "package  main\n\nvar  i int",
		out: "REPL 1 3\npackage main\n\nvar i int",
	},
}

func TestFormat(t *testing.T) {
	for _, tt := range formatTests {
		var buf bytes.Buffer
		doFormat(&buf, strings.NewReader(tt.in), []string{"test.go"}, tt.fmtonly)
		out := buf.String()
		if out != tt.out {
			t.Errorf("%q: got %q, want %q", tt.in, out, tt.out)
		}
	}
}
