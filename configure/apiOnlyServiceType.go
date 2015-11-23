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

type ApiOnlyServiceType struct {
	name string
}

func newApiOnlyServiceType() ApiOnlyServiceType {
	return ApiOnlyServiceType{name: "empty service"}
}

func (s ApiOnlyServiceType) Name() string {
	return s.name
}

func (s ApiOnlyServiceType) Serve() {
	listenForHttp()
}

func (s ApiOnlyServiceType) Configure() {
	createApiEndpoint()
}

func listenForHttp() {
	http.ListenAndServe(":9000", nil)
}

func createApiEndpoint() {
	http.HandleFunc("/resemble", api)
}

func api(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		"API endpoint",
	)
}
