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
	"log"
	"net/http"
)

type HttpServiceType struct {
	name      string
	endpoints []HttpEndpoint
}

func newHttpServiceType(config ResembleConfig) HttpServiceType {
	endpoints, err := config.createServiceFromConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return HttpServiceType{name: "HTTP", endpoints: endpoints}
}

func (s HttpServiceType) Name() string {
	return s.name
}

func (s HttpServiceType) Serve() {
	listenForHttp()
}

func (s HttpServiceType) Configure() {
	createApiEndpoint()
	http.HandleFunc("/", HandleHttp(s.endpoints))
}

func HandleHttp(endpoints []HttpEndpoint) http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {
		if matched, endpoint := matchHttpRequest(endpoints, req); matched {
			endpoint.Respond(res, req)
		} else {
			http.NotFound(res, req)
		}
	}
}

func matchHttpRequest(endpoints []HttpEndpoint, req *http.Request) (matched bool, endpoint HttpEndpoint) {
	if len(endpoints) == 0 {
		return false, NamedHttpEndpoint{}
	}
	for i := 0; i < len(endpoints); i++ {
		if endpoints[i].Match(req) {
			return true, endpoints[i]
		} else {
			continue
		}
	}
	return false, NamedHttpEndpoint{}
}
