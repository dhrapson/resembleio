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
	"errors"
	"gopkg.in/yaml.v2"
	"strings"
)

func ConfigureService(configYaml string) (s ServiceType, err error) {
	if len(configYaml) == 0 {
		return newApiOnlyServiceType(), err
	}

	if strings.Contains(configYaml, "REST_HTTP") {
		s = newRestServiceType()
	} else if strings.Contains(configYaml, "HTTP") {
		var resembleConfig ResembleConfig
		resembleConfig, err = ReadHttpConfiguration(configYaml)
		s = newHttpServiceType(resembleConfig)
	} else {
		err = errors.New("Invalid config: `type`")
	}
	return s, err
}

func ReadHttpConfiguration(configYaml string) (resembleConfig ResembleConfig, err error) {
	resembleConfig = ResembleConfig{}

	err = yaml.Unmarshal([]byte(configYaml), &resembleConfig)
	if err != nil {
		return resembleConfig, errors.New("Error reading YAML text: " + err.Error())
	}
	return resembleConfig, err
}
