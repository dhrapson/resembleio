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
	"strings"
)

func ConfigureService(configYaml string) ServiceType {
	var s ServiceType
	if len(configYaml) == 0 {
		s = newApiOnlyServiceType()
	} else if strings.Contains(configYaml, "type: REST_HTTP") {
		s = newRestServiceType()
	} else if strings.Contains(configYaml, "type: HTTP") {
		s = newHttpServiceType()
	} else {
		s = newHttpServiceType()
	}
	return s
}
