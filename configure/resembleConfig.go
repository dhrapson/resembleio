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
	"log"
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
	if _, map_err := yaml.Map(); map_err != nil {
		return errors.New("Error reading YAML text")
	}

	typeName, err := yaml.Get("type").String()
	if err != nil {
		return errors.New("Resemble config: invalid `type`")
	}
	c.TypeName = typeName

	matchersYaml := yaml.Get("matchers")
	if err == nil {
		c.Matchers, err = getMatchers(matchersYaml)
	}
	return nil
}

func getMatchers(matchersYaml *simpleyaml.Yaml) (matchers []HttpMatcher, err error) {

	matchersType, err := matchersYaml.Array()
	arraySize := len(matchersType)
	matchers = make([]HttpMatcher, arraySize)
	for i := 0; i < arraySize; i++ {
		matcher := matchersYaml.GetIndex(i)
		matcherType, err := matcher.Get("type").String()
		if err != nil {
			return matchers, errors.New("Resemble config: invalid `matcher->type`")
		}
		if matcherType == "url_path" {
			path_regex, err := matcher.Get("path_regex").String()
			if err != nil {
				log.Fatalln("missing path_regex for url_path matcher")
			}
			matchers[i], err = NewUrlPathHttpMatcher(path_regex)
			if err != nil {
				log.Fatalln("invalid path_regex for url_path matcher", path_regex)
			}
		} else {
			log.Fatalln("missing url_path for ")
		}
	}
	return matchers, err
}
