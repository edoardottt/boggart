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

package template

import "errors"

var (
	ErrMissingType                = errors.New("template: missing template type")
	ErrUniqueRequestID            = errors.New("template: request IDs are not unique")
	ErrUniqueRequestEndpoint      = errors.New("template: request endpoints are not unique")
	ErrMissingDefaultRequest      = errors.New("template: missing default request")
	ErrMandatoryIP                = errors.New("template: ip is mandatory")
	ErrMissingID                  = errors.New("template: missing id in request")
	ErrMissingDefaultResponseType = errors.New("template: missing response type in default request")
	ErrMissingDefaultContentType  = errors.New("template: missing content type in default request")
	ErrMissingDefaultContent      = errors.New("template: missing content in default request")
	ErrDuplicatePathsIgnore       = errors.New("template: duplicate paths in ignore array")
	ErrMissingSlashIgnore         = errors.New("template: all paths in ignore array must start with a forward slash")
	ErrPathRequestIgnore          = errors.New("template: path defined both in ignore and requests")
	ErrMissingEndpointID          = errors.New("template: missing endpoint in request with id")
	ErrMissingMethodsID           = errors.New("template: missing methods in request with id")
	ErrMissingResponseTypeID      = errors.New("template: missing response type in request with id")
	ErrMissingContentTypeID       = errors.New("template: missing content type in request with id")
	ErrMissingContentID           = errors.New("template: missing content in request with id")
)
