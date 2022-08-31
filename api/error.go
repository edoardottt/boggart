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

package api

import "errors"

var (
	ErrPossibleTopValue  = errors.New("possible values for top: method / path / body")
	ErrNoStatsIP         = errors.New("no stats available for the specified IP")
	ErrIDDefined         = errors.New("if id is defined, no other parameters need to be defined")
	ErrDateDefined       = errors.New("if date is defined, lt and gt must be blank")
	ErrLtBeforeGt        = errors.New("lt cannot be before gt")
	ErrHTTPMethodUnknown = errors.New("http method unknown")
	ErrInvalidIP         = errors.New("ip address is not valid")
	ErrDatetimeFormat    = errors.New("correct datetime format: 2006-01-02T15:04:05-0700")
)
