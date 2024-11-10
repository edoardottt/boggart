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

package time

import "time"

// TranslateTime.
func TranslateTime(input string) (time.Time, error) {
	t, err := time.Parse("2006-01-02T15:04:05-0700", input)
	if err != nil {
		return time.Time{}, ErrDatetimeFullFormat
	}

	return t, nil
}

// GetDay.
func GetDay(input string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", input)
	if err != nil {
		return time.Time{}, ErrDatetimeDayFormat
	}

	return t, nil
}
