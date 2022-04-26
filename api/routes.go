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

//Api routes
const (
	IPInfo      = "/api/{ip}"
	LogsIP      = "/api/logs/ip/{ip}"
	LogsIPDate  = "/api/logs/ip/{ip}/{date}"
	LogsIPRange = "/api/logs/ip/{ip}/{range}"

	LogsPath      = "/api/logs/path/{path}"
	LogsPathDate  = "/api/logs/path/{path}/{date}"
	LogsPathRange = "/api/logs/path/{path}/{range}"

	LogsMethod      = "/api/logs/method/{method}"
	LogsMethodDate  = "/api/logs/method/{method}/{date}"
	LogsMethodRange = "/api/logs/method/{method}/{range}"
)
