/*
=======================
	boggart
=======================

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:	https://github.com/edoardottt/boggart
	@Author:		edoardottt, https://www.edoardoottavianelli.it
	@License:		https://github.com/edoardottt/boggart/blob/main/LICENSE
*/

package slice_test

import (
	"testing"

	"github.com/edoardottt/boggart/internal/slice"
	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicateValues(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "empty",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "123",
			input: []string{"1", "2", "3"},
			want:  []string{"1", "2", "3"},
		},
		{
			name:  "123dups",
			input: []string{"1", "2", "3", "1", "2", "3", "1", "2", "3"},
			want:  []string{"1", "2", "3"},
		},
		{
			name:  "aaa",
			input: []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"},
			want:  []string{"a"},
		},
		{
			name: "alphabet",
			input: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
				"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
				"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
				"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
			want: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
				"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.RemoveDuplicateValues(tt.input)
			require.Equal(t, got, tt.want)
		})
	}
}

func TestStringInSlice(t *testing.T) {
	tests := []struct {
		name   string
		inputA string
		inputB []string
		want   bool
	}{
		{
			name:   "true",
			inputA: "a",
			inputB: []string{"a", "b"},
			want:   true,
		},
		{
			name:   "false",
			inputA: "a",
			inputB: []string{"b", "c"},
			want:   false,
		},
		{
			name:   "empty",
			inputA: "a",
			inputB: []string{},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice.StringInSlice(tt.inputA, tt.inputB)
			require.Equal(t, got, tt.want)
		})
	}
}
