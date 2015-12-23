//Copyright (C) 2015 dhrapson

//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
//along with this program.  If not, see <http://www.gnu.org/licenses/>.

package configure

import (
	"net/http"
	"regexp"
)

type KeyValuesHttpMatcher struct {
	KeyRegex   *regexp.Regexp
	ValueRegex *regexp.Regexp
}

func (m KeyValuesHttpMatcher) Match(req *http.Request) bool {
	return m.MatchKeyValues(req.URL.Query())
}

func (m KeyValuesHttpMatcher) MatchKeyValues(params map[string][]string) bool {
	for key, val := range params {
		if m.KeyRegex.MatchString(key) {
			for i := 0; i < len(val); i++ {
				if m.ValueRegex.MatchString(val[i]) {
					return true
				}
			}
		}
	}
	return false
}
