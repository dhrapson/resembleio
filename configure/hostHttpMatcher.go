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

type HostHttpMatcher struct {
	regex *regexp.Regexp
}

func NewHostHttpMatcher(regularExpr string) (matcher HostHttpMatcher, err error) {
	compiledRegex, err := regexp.Compile(regularExpr)
	if err == nil {
		matcher.regex = compiledRegex
	}
	return matcher, err
}

func (m HostHttpMatcher) Match(req *http.Request) bool {
	return m.MatchHost(req.Host)
}

func (m HostHttpMatcher) MatchHost(host string) bool {
	return m.regex.MatchString(host)
}
