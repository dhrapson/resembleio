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
	"github.com/smallfish/simpleyaml"
)

type ResembleConfig struct {
	TypeName string `yaml:"type"`
	Matchers []HttpMatcher
}

func (c *ResembleConfig) Parse(data []byte) error {

	yaml, err := simpleyaml.NewYaml(data)
	if err != nil {
		return errors.New("Error reading YAML text")
	}
	// Probably a defect in the simpleyaml library, the NewYaml function usage above isnt throwing an error
	// when we get it to open a totally invalid yaml file. Dealing with that separately here....
	if _, mapErr := yaml.Map(); mapErr != nil {
		return errors.New("Error reading YAML text")
	}

	typeName, err := yaml.Get("type").String()
	if err != nil {
		return errors.New("Resemble config: invalid `type`")
	}
	c.TypeName = typeName

	matchersYaml := yaml.Get("matchers")
	if err == nil {
		c.Matchers, err = getMatchersFromYaml(matchersYaml)
	}
	return nil
}

func getMatchersFromYaml(matchersYaml *simpleyaml.Yaml) (matchers []HttpMatcher, err error) {

	matchersMap, mapErr := matchersYaml.Map()
	if mapErr != nil {
		return matchers, errors.New("Error reading matchers node")
	}
	arraySize := len(matchersMap)
	matchers = make([]HttpMatcher, arraySize)

	var count int
	path_regex, err := matchersYaml.Get("path_regex").String()
	if err == nil {
		matchers[count], err = NewUrlPathHttpMatcher(path_regex)
		if err != nil {
			return matchers, err
		}
		count++
	}

	verb_regex, err := matchersYaml.Get("verb_regex").String()
	if err == nil {
		matchers[count], err = NewHttpVerbMatcher(verb_regex)
		if err != nil {
			return matchers, err
		}
		count++
	}
	return matchers, err
}
