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
	@Author:		edoardottt, https://edoardottt.com
	@License:		https://github.com/edoardottt/boggart/blob/main/LICENSE
*/

package template

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckShodanTemplate(t *testing.T) {
	tests := []struct {
		name    string
		tmpl    Template
		wantErr error
	}{
		{
			name: "valid shodan template with ip",
			tmpl: Template{
				Type: ShodanTemplateType,
				IP:   "1.2.3.4",
			},
			wantErr: nil,
		},
		{
			name: "shodan template missing ip",
			tmpl: Template{
				Type: ShodanTemplateType,
				IP:   "",
			},
			wantErr: ErrMandatoryIP,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckShodanTemplate(tt.tmpl)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestCheckTemplateShodan(t *testing.T) {
	tests := []struct {
		name    string
		tmpl    Template
		wantErr error
	}{
		{
			name: "shodan type with ip passes",
			tmpl: Template{
				Type: ShodanTemplateType,
				IP:   "8.8.8.8",
			},
			wantErr: nil,
		},
		{
			name: "shodan type without ip fails",
			tmpl: Template{
				Type: ShodanTemplateType,
			},
			wantErr: ErrMandatoryIP,
		},
		{
			name:    "missing type fails",
			tmpl:    Template{},
			wantErr: ErrMissingType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckTemplate(tt.tmpl)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestIDUniqueShodan(t *testing.T) {
	tmpl := Template{
		Type:     ShodanTemplateType,
		IP:       "1.2.3.4",
		Requests: []Request{{ID: "a"}, {ID: "a"}},
	}
	require.True(t, IDUnique(tmpl))
}

func TestEndpointUniqueShodan(t *testing.T) {
	tmpl := Template{
		Type:     ShodanTemplateType,
		IP:       "1.2.3.4",
		Requests: []Request{{Endpoint: "/x"}, {Endpoint: "/x"}},
	}
	require.True(t, EndpointUnique(tmpl))
}

func TestMissingTemplateDefaultShodan(t *testing.T) {
	tmpl := Template{
		Type: ShodanTemplateType,
		IP:   "1.2.3.4",
	}
	require.False(t, MissingTemplateDefault(tmpl))
}

func TestDefaultShodan(t *testing.T) {
	tmpl := Template{
		Type: ShodanTemplateType,
		IP:   "1.2.3.4",
		Requests: []Request{
			{ID: DefaultID, Content: "hello"},
		},
	}
	require.Equal(t, Request{}, Default(tmpl))
}

func TestHTTPMethodsAsString(t *testing.T) {
	tests := []struct {
		name    string
		methods []HTTPMethod
		want    []string
	}{
		{
			name:    "empty",
			methods: []HTTPMethod{},
			want:    []string{},
		},
		{
			name:    "single",
			methods: []HTTPMethod{MethodGet},
			want:    []string{"GET"},
		},
		{
			name:    "multiple",
			methods: []HTTPMethod{MethodGet, MethodPost, MethodDelete},
			want:    []string{"GET", "POST", "DELETE"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HTTPMethodsAsString(tt.methods)
			require.Equal(t, tt.want, got)
		})
	}
}
