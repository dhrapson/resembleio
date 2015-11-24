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

type QueryParamHttpMatcher struct {
	keyRegex   *regexp.Regexp
	valueRegex *regexp.Regexp
}

func NewQueryParamHttpMatcher(keyRegex string, valueRegex string) (matcher QueryParamHttpMatcher, err error) {

	compiledKeyRegex, err := regexp.Compile(keyRegex)
	if err != nil {
		return matcher, err
	}
	matcher.keyRegex = compiledKeyRegex
	compiledValueRegex, err := regexp.Compile(valueRegex)
	if err == nil {
		matcher.valueRegex = compiledValueRegex
	}
	return matcher, err
}

func (m QueryParamHttpMatcher) Match(req *http.Request) bool {
	return m.MatchParam(req.URL.Query())
}

func (m QueryParamHttpMatcher) MatchParam(params map[string][]string) bool {
	for key, val := range params {
		if m.keyRegex.MatchString(key) {
			for i := 0; i < len(val); i++ {
				if m.valueRegex.MatchString(val[i]) {
					return true
				}
			}
		}
	}
	return false
}
