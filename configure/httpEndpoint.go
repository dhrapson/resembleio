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
)

type HttpEndpoint interface {
	Match(req *http.Request) bool
}

type NamedHttpEndpoint struct {
	Name string
  Matchers []HttpMatcher
}

func (e NamedHttpEndpoint) Match(req *http.Request) bool {
	if len(e.Matchers) == 0 {
		return false
	}
	for i := 0; i < len(e.Matchers); i++ {
		if e.Matchers[i].Match(req) {
			return true
		} else {
			continue
		}
	}
	return false
}
