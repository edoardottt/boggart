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
	MissingTypeErr                = errors.New("template: missing template type")
	UniqueRequestIDErr            = errors.New("template: request IDs are not unique")
	UniqueRequestEndpointErr      = errors.New("template: request endpoints are not unique")
	MissingDefaultRequestErr      = errors.New("template: missing default request")
	MandatoryIPErr                = errors.New("template: ip is mandatory")
	MissingIDErr                  = errors.New("template: missing id in request")
	MissingDefaultResponseTypeErr = errors.New("template: missing response type in default request")
	MissingDefaultContentTypeErr  = errors.New("template: missing content type in default request")
	MissingDefaultContentErr      = errors.New("template: missing content in default request")
	DuplicatePathsIgnoreErr       = errors.New("template: duplicate paths in ignore array")
	MissingSlashIgnoreErr         = errors.New("template: all paths in ignore array must start with a forward slash")
	PathRequestIgnoreErr          = errors.New("template: path defined both in ignore and requests")
	MissingEndpointIDErr          = errors.New("template: missing endpoint in request with id ")
	MissingMethodsIDErr           = errors.New("template: missing methods in request with id ")
	MissingResponseTypeIDErr      = errors.New("template: missing response type in request with id ")
	MissingContentTypeIDErr       = errors.New("template: missing content type in request with id ")
	MissingContentIDErr           = errors.New("template: missing content in request with id ")
)
