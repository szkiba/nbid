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
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getopt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args []string
		want *options
	}{
		{
			name: "defaults",
			want: &options{},
		},
		{
			name: "version",
			want: &options{version: true},
			args: []string{"-v"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, getopt(append([]string{"nbid"}, tt.args...)))
		})
	}
}

func Test_getid(t *testing.T) {
	t.Parallel()

	input := "The quick brown fox jumps over the lazy dog"
	assert.Equal(t, "QUKFNCO7QU098QEAJAUB021E9S", getid(input))

	id1 := getid("")
	id2 := getid("")

	assert.NotEmpty(t, id1)
	assert.NotEmpty(t, id2)
	assert.NotEqual(t, id1, id2)
}

func Test_getver(t *testing.T) {
	t.Parallel()

	v := getver()

	assert.True(t, strings.HasPrefix(v, "nbid"))
	assert.True(t, strings.HasSuffix(v, fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)))
}
