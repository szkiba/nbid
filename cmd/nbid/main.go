// MIT License
//
// Copyright (c) 2021 Iván Szkiba
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/szkiba/nbid"
)

var version = "dev"

type options struct {
	version bool
	input   string
}

const usage = `usage: %s [name]

Generate NBID for name, or random NBID if name is missing.

Example: %s "The quick brown fox jumps over the lazy dog"
Output: QUKFNCO7QU098QEAJAUB021E9S

`

func getopt(args []string) *options {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	flags.Usage = func() {
		fmt.Fprintf(flags.Output(), usage, flags.Name(), flags.Name())
		flags.PrintDefaults()
	}

	o := options{}

	ver := flags.Bool("v", false, "prints version")

	_ = flags.Parse(args[1:])

	o.version = *ver
	o.input = flags.Arg(0)

	return &o
}

func getid(s string) string {
	if s == "" {
		return nbid.Random().String()
	}

	return nbid.New([]byte(s)).String()
}

func getver() string {
	return fmt.Sprintf("nbid/%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
}

func main() {
	o := getopt(os.Args)

	if o.version {
		fmt.Fprintln(os.Stderr, getver())
	} else {
		fmt.Fprintln(os.Stdout, getid(o.input))
	}
}
