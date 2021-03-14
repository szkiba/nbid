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

package nbid

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"hash"
)

var (
	// ErrInvalidID is returned when trying to unmarshal an invalid NBID.
	ErrInvalidID = errors.New("nbid: invalid NBID")

	// Nil is an empty NBID (all zero).
	Nil NBID
)

const (
	encodedLen = 26 // string encoded len
	rawLen     = 16 // binary raw len
)

// NBID represents a Name Based ID.
// The binary representation of the NBID is a 16 byte byte array.
// The string representation is using base32 hex (w/o padding).
type NBID [rawLen]byte

// NewHash returns a new NBID derived from the hash of data generated by h.
// The hash should be at least 16 byte in length.
// The first 16 bytes of the hash are used to form the ID.
func NewHash(h hash.Hash, data []byte) NBID {
	h.Reset()
	h.Write(data) //nolint:errcheck
	s := h.Sum(nil)

	var id NBID

	copy(id[:], s)

	return id
}

// New returns a new NBID derived from the SHA256 hash of data.
// The first 16 bytes of the SHA256 hash are used to form the ID.
// It is the same as calling:
//
//  NewHash(sha256.New(), data)
func New(data []byte) NBID {
	return NewHash(sha256.New(), data)
}

// Random returns a new random generated NBID.
// The strength of the IDs is based on the strength of the crypto/rand
// package.
func Random() NBID {
	var id NBID

	rand.Read(id[:]) //nolint:errcheck

	return id
}

// Parse decodes s into an NBID or returns an error.
// The string representation is using base32 hex (w/o padding).
// Returns an error if the s does not have a length of 16.
func Parse(s string) (NBID, error) {
	i := &NBID{}
	err := i.UnmarshalText([]byte(s))

	return *i, err
}

// String returns the string form of NBID, a base32 hex (w/o padding).
// RFC 4648 / 7. Base 32 Encoding with Extended Hex Alphabet
// https://tools.ietf.org/html/rfc4648#section-7
func (id NBID) String() string {
	return base32.HexEncoding.WithPadding(base32.NoPadding).EncodeToString(id[:])
}

// MustParse decodes s into an NBID or panics if the string cannot be parsed.
func MustParse(s string) NBID {
	id, err := Parse(s)
	if err != nil {
		panic(err)
	}

	return id
}

// MarshalText implements encoding/text TextMarshaler interface.
func (id NBID) MarshalText() (text []byte, err error) {
	return []byte(id.String()), nil
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (id NBID) MarshalBinary() ([]byte, error) {
	return id[:], nil
}

// MarshalJSON implements encoding/json Marshaler interface.
func (id NBID) MarshalJSON() ([]byte, error) {
	if id.IsNil() {
		return []byte("null"), nil
	}

	text, err := id.MarshalText()

	return []byte(`"` + string(text) + `"`), err
}

// UnmarshalText implements encoding/text TextUnmarshaler interface.
func (id *NBID) UnmarshalText(text []byte) error {
	if len(text) != encodedLen {
		return ErrInvalidID
	}

	_, err := base32.HexEncoding.WithPadding(base32.NoPadding).Decode(id[:], text)
	if err != nil {
		return ErrInvalidID
	}

	return nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (id *NBID) UnmarshalBinary(data []byte) error {
	if len(data) != rawLen {
		return ErrInvalidID
	}

	copy(id[:], data)

	return nil
}

// UnmarshalJSON implements encoding/json Unmarshaler interface.
func (id *NBID) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		*id = Nil

		return nil
	}

	return id.UnmarshalText(b[1 : len(b)-1])
}

// IsNil returns true if this is a "nil" NBID.
func (id NBID) IsNil() bool {
	return id.Equal(Nil)
}

// Bytes returns the byte array representation of NBID.
func (id NBID) Bytes() []byte {
	return id[:]
}

// FromBytes creates a new NBID from a byte slice.
// Returns an error if the slice does not have a length of 16.
// The bytes are copied from the slice.
func FromBytes(b []byte) (NBID, error) {
	var id NBID

	if len(b) != rawLen {
		return id, ErrInvalidID
	}

	copy(id[:], b)

	return id, nil
}

// Compare returns an integer comparing two NBIDs. It behaves just like `bytes.Compare`.
func (id NBID) Compare(other NBID) int {
	return bytes.Compare(id[:], other[:])
}

// Equal returns true if two NBID is equal. It behaves just like `bytes.Equal`.
func (id NBID) Equal(other NBID) bool {
	return bytes.Equal(id[:], other[:])
}
