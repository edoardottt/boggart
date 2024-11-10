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

package api

// Api routes.
const (
	// This is the 'fallback' endpoint, when
	// a client tries to request a resource on an
	// endpoint that is not defined, this will be
	// the default response it will get.
	NotFound = "/"

	// This is the health status endpoint.
	// Performing a request to this endpoint the
	// client should receive a lightweight response
	// 'OK' with 200 as status code if everything is
	// behaving correctly.
	Health = "/api/health"

	// This endpoint is intended to serve a summary
	// of the information available for the IP taken
	// as input.
	// Parameter: top
	// Which info? Number of requests, timestamp of
	// the last activity, top X (default 10)
	// methods, path, headers.
	IPInfo = "/api/info/{ip}"

	// This endpoint is intended to serve information about
	// logs, so the requests were made to the Honeypot.
	// These are the parameters the endpoint accepts:
	// - id
	// - ip
	// - method
	// - header
	// - path
	// - date (YYYY-MM-DD)
	// - lt (less than YYYY-MM-DD-HH-MM-SS)
	// - gt (greater than YYYY-MM-DD-HH-MM-SS).
	APILogs = "api/logs"

	// This endpoint is intended to perform a heavy
	// and accurate scan on the logs. It takes as input
	// these parameters:
	// - regex (Go format)
	// - attack (use a list of well known regex)
	// - target (where to apply the regex)
	// - ip
	// - method
	// - header
	// - path
	// - date
	// - lt (less than YYYY-MM-DD-HH-MM-SS)
	// - gt (greater than YYYY-MM-DD-HH-MM-SS).
	APIDetect = "api/detect"

	// This endpoint gives a general overview of the system.
	APIStats = "api/stats"

	// This endpoint gives a detailed overview of the data
	// stored in the DB.
	APIStatsDB = "api/stats/db"
)
