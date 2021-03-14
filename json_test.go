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
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szkiba/nbid"
)

func TestJSON(t *testing.T) {
	t.Parallel()

	type x struct {
		ID nbid.NBID `json:"id"`
	}

	tests := []struct {
		name string
		json string
		id   string
	}{
		{
			name: "normal",
			json: `{"id":"ABCDEFGHIJKLMNOPQRSTUV1234"}`,
			id:   "ABCDEFGHIJKLMNOPQRSTUV1234",
		},
		{
			name: "null",
			json: `{"id":null}`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var id nbid.NBID

			if tt.id != "" {
				assert.Nil(t, id.UnmarshalText([]byte(tt.id)))
			}

			data := x{}
			err := json.Unmarshal([]byte(tt.json), &data)

			assert.Nil(t, err)
			assert.Equal(t, id, data.ID)

			b, err := json.Marshal(data)
			assert.Nil(t, err)

			assert.Equal(t, tt.json, string(b))
		})
	}
}
