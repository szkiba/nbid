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
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/szkiba/nbid"

	"github.com/stretchr/testify/assert"
)

func TestMustParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		in        string
		out       string
		wantPanic bool
	}{
		{name: "normal", in: "ABCDEFGHIJKLMNOPQRSTUV1234", out: "ABCDEFGHIJKLMNOPQRSTUV1234"},
		{name: "error_size", in: "small", wantPanic: true},
		{name: "error_value", in: "XXXXXXXXXXXXXXXXXXXXXXXXXX", wantPanic: true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			id := nbid.MustParse(tt.in)

			txt, err := id.MarshalText()
			assert.Nil(t, err)
			assert.Equal(t, tt.out, string(txt))
		})
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	data := []byte("The quick brown fox jumps over the lazy dog")

	id := nbid.New(data)
	idHash := nbid.NewHash(sha256.New(), data)

	assert.NotEqual(t, nbid.Nil, id)
	assert.Equal(t, idHash, id)

	assert.Equal(t, "QUKFNCO7QU098QEAJAUB021E9S", id.String())
}

func TestRandom(t *testing.T) {
	t.Parallel()

	id1 := nbid.Random()
	id2 := nbid.Random()

	assert.NotEqual(t, nbid.Nil, id1)
	assert.NotEqual(t, nbid.Nil, id2)
	assert.NotEqual(t, id1, id2)
}

func TestIsNil(t *testing.T) {
	t.Parallel()

	var id nbid.NBID

	assert.True(t, id.IsNil())
	assert.False(t, nbid.Random().IsNil())
}

func TestCompareAndEqual(t *testing.T) {
	t.Parallel()

	a := nbid.NBID{1}
	b := nbid.NBID{2}
	c := nbid.NBID{1}

	assert.True(t, a.Equal(c))
	assert.False(t, a.Equal(b))

	assert.Less(t, a.Compare(b), 0)
	assert.Greater(t, b.Compare(a), 0)
	assert.Equal(t, a.Compare(c), 0)
}

func TestText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		in      string
		out     string
		wantErr bool
	}{
		{name: "normal", in: "ABCDEFGHIJKLMNOPQRSTUV1234", out: "ABCDEFGHIJKLMNOPQRSTUV1234"},
		{name: "zero", in: "AAAAAAAAAAAAAAAAAAAAAAAAA8", out: "AAAAAAAAAAAAAAAAAAAAAAAAA8"},
		{name: "error_size", in: "small", wantErr: true},
		{name: "error_value", in: "XXXXXXXXXXXXXXXXXXXXXXXXXX", wantErr: true},
		{name: "sample1", in: "DEJRG44TLK8T305K0304VL1GP0", out: "DEJRG44TLK8T305K0304VL1GP0"},
		{name: "sample2", in: "DEJRG4CTLK8T305K0304VL1GP0", out: "DEJRG4CTLK8T305K0304VL1GP0"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			id, err := nbid.Parse(tt.in)

			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			txt, err := id.MarshalText()
			assert.Nil(t, err)
			assert.Equal(t, tt.out, string(txt))
		})
	}
}

func TestBinary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		hex     string
		wantErr bool
	}{
		{name: "normal", hex: "e12340c084a111ebb6a7af57dd3dfeda"},
		{name: "zero", hex: "00000000000000000000000000000000"},
		{name: "error_size", hex: "42", wantErr: true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b, err := hex.DecodeString(tt.hex)

			assert.Nil(t, err)

			id1, err1 := nbid.FromBytes(b)

			var id2 nbid.NBID
			err2 := id2.UnmarshalBinary(b)

			if tt.wantErr {
				assert.Error(t, err1)
				assert.Error(t, err2)

				return
			}

			// FromBytes and UnmarshalBinary should result same id
			assert.Equal(t, b, id1[:])
			assert.Equal(t, b, id2[:])

			assert.Equal(t, b, id1.Bytes())
			assert.Equal(t, b, id2.Bytes())

			// MarshalBinary should result same byte array
			b1, err := id1.MarshalBinary()
			assert.Nil(t, err)

			b2 := id1.Bytes()

			assert.Equal(t, b, b1)
			assert.Equal(t, b, b2)

			assert.True(t, id1.Equal(id2))
			assert.True(t, id2.Equal(id1))
		})
	}
}
