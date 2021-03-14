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

package nbid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szkiba/nbid"
)

func TestScan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		src     interface{}
		want    nbid.NBID
		wantErr bool
	}{
		{
			name: "normal",
			src:  "QUKFNCO7QU098QEAJAUB021E9S",
			want: nbid.MustParse("QUKFNCO7QU098QEAJAUB021E9S"),
		},
		{name: "nil", src: nil, want: nbid.Nil},
		{name: "empty_string", src: "", want: nbid.Nil},
		{
			name:    "invalid_string",
			src:     "XXX",
			wantErr: true,
		},
		{name: "empty_bytes", src: []byte{}, want: nbid.Nil},
		{
			name: "bytes_string",
			src:  []byte("QUKFNCO7QU098QEAJAUB021E9S"),
			want: nbid.MustParse("QUKFNCO7QU098QEAJAUB021E9S"),
		},
		{
			name: "bytes",
			src:  nbid.MustParse("QUKFNCO7QU098QEAJAUB021E9S").Bytes(),
			want: nbid.MustParse("QUKFNCO7QU098QEAJAUB021E9S"),
		},
		{name: "invalid_type", src: 42, wantErr: true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var id nbid.NBID

			err := id.Scan(tt.src)

			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tt.want, id)
		})
	}
}

func TestValue(t *testing.T) {
	t.Parallel()

	id := nbid.Random()

	val, err := id.Value()
	assert.Nil(t, err)
	assert.Equal(t, id.String(), val)
}
