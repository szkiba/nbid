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
	"database/sql/driver"
	"fmt"
	"reflect"
)

var invalidValue = reflect.Value{}

// Scan implements sql.Scanner so NBIDs can be read from databases transparently.
// Currently, database types that map to string and []byte are supported.
func (id *NBID) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil

	case string:
		if src == "" {
			return nil
		}

		u, err := Parse(src)
		if err != nil {
			return err
		}

		*id = u

	case []byte:
		if len(src) == 0 {
			return nil
		}

		if len(src) != rawLen {
			return id.Scan(string(src))
		}

		copy((*id)[:], src)

	default:
		return fmt.Errorf("%w: unable to scan type %T", ErrInvalidID, src)
	}

	return nil
}

// Value implements sql.Valuer so that NBIDs can be written to databases transparently.
func (id NBID) Value() (driver.Value, error) {
	return id.String(), nil
}
