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

package net_test

import (
	"testing"

	"github.com/edoardottt/boggart/internal/net"
	"github.com/stretchr/testify/require"
)

func TestValidIPAddress(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "empty input",
			input: "",
			want:  false,
		},
		{
			name:  "true",
			input: "1.1.1.1",
			want:  true,
		},
		{
			name:  "false",
			input: "2i7vbyyvb28qytv82ot",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := net.ValidIPAddress(tt.input)
			require.Equal(t, got, tt.want)
		})
	}
}

func TestPrivateIP(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "trueA",
			input: "10.1.1.1",
			want:  true,
		},
		{
			name:  "trueB",
			input: "172.18.1.1",
			want:  true,
		},
		{
			name:  "trueC",
			input: "192.168.1.1",
			want:  true,
		},
		{
			name:  "false",
			input: "11.56.3.9",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := net.PrivateIP(tt.input)
			require.Nil(t, err)
			require.Equal(t, got, tt.want)
		})
	}
}
