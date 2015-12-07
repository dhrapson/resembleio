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
	"errors"
)

type KeyValuesHttpMatcher struct {
  KeyRegexString string `yaml:"key_regex"`
  ValueRegexString string `yaml:"value_regex"`
  keyRegex *regexp.Regexp
  valueRegex *regexp.Regexp
}

func NewKeyValuesHttpMatcher(keyRegex string, valueRegex string) (matcher KeyValuesHttpMatcher, err error) {

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

func (m *KeyValuesHttpMatcher) acceptValidation() (err error) {
	m.keyRegex, err = regexp.Compile(m.KeyRegexString)
	if err != nil {
		return errors.New("Error validating YAML text: " + err.Error())
	}
	m.valueRegex, err = regexp.Compile(m.ValueRegexString)
	if err != nil {
		return errors.New("Error validating YAML text: " + err.Error())
	}
	return nil
}

func (m KeyValuesHttpMatcher) Match(req *http.Request) bool {
	return m.MatchKeyValues(req.URL.Query())
}

func (m KeyValuesHttpMatcher) MatchKeyValues(params map[string][]string) bool {
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
