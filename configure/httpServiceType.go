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
	"io"
	"net/http"
)

type HttpServiceType struct {
	name     string
	matchers []HttpMatcher
}

func newHttpServiceType() HttpServiceType {
	m := []HttpMatcher{}
	return HttpServiceType{name: "HTTP", matchers: m}
}

func (s HttpServiceType) Name() string {
	return s.name
}

func (s HttpServiceType) Serve() {
	listenForHttp()
}

func (s HttpServiceType) Configure() {
	createApiEndpoint()
	http.HandleFunc("/", HttpEndpoint(s.matchers))
}

func HttpEndpoint(matchers []HttpMatcher) http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {
		if matchHttpRequest(matchers, req) {
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			io.WriteString(
				res,
				"HTTP endpoint - matched",
			)
		} else {
			http.NotFound(res, req)
		}
	}
}

func matchHttpRequest(matchers []HttpMatcher, req *http.Request) bool {
	if len(matchers) == 0 {
		return false
	}
	for i := 0; i < len(matchers); i++ {
		if matchers[i].Match(req) {
			return true
		}
	}
	return false
}
