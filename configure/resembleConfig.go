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
	"regexp"
)

type ResembleConfig struct {
	TypeName        string               `yaml:"type"`
	EndpointConfigs []HttpEndpointConfig `yaml:"endpoints"`
	ModeConfigs     []ModeConfig         `yaml:"modes"`
}

type ModeConfig struct {
	Name string
}

type HttpEndpointConfig struct {
	Name             string
	MatcherConfigs   []HttpRequestMatcherConfig `yaml:"matchers"`
	ResponderConfigs []HttpResponderConfig      `yaml:"responders"`
}

type HttpRequestMatcherConfig struct {
	Name            string
	VerbRegexString string                       `yaml:"verb_regex"`
	HostRegexString string                       `yaml:"host_regex"`
	PathRegexString string                       `yaml:"path_regex"`
	QueryParams     []KeyValuesHttpMatcherConfig `yaml:"query_params"`
	Headers         []KeyValuesHttpMatcherConfig `yaml:"headers"`
}

type KeyValuesHttpMatcherConfig struct {
	KeyRegexString   string `yaml:"key_regex"`
	ValueRegexString string `yaml:"value_regex"`
}

type HttpResponderConfig struct {
	Name    string
	Mode    string
	Content HttpResponderContent
}

type HttpResponderContent struct {
	ContentType string `yaml:"type"`
	SourceDir   string `yaml:"source_dir"`
	TmpDir      string `yaml:"tmp_dir"`
	Script      string
	SourceFile  string `yaml:"source_file"`
}

func (config ResembleConfig) createServiceFromConfig() (endpoints []HttpEndpoint, err error) {
	endpoints = []HttpEndpoint{}
	for _, endpoint := range config.EndpointConfigs {
		newEndpoint := NamedHttpEndpoint{endpoint.Name, []HttpMatcher{}}
		newEndpoint.Matchers, err = endpoint.createMatchersFromConfig()
		if err != nil {
			return nil, errors.New("Error validating YAML text: " + err.Error())
		}
		endpoints = append(endpoints, newEndpoint)
	}
	return endpoints, err
}

func (config HttpEndpointConfig) createMatchersFromConfig() (matchers []HttpMatcher, err error) {
	matchers = []HttpMatcher{}
	for _, matcher := range config.MatcherConfigs {
		newMatcher, err := matcher.NewMatcher()
		if err != nil {
			return nil, errors.New("Error validating YAML text: " + err.Error())
		}
		matchers = append(matchers, newMatcher)
	}
	return matchers, err
}

func (m HttpRequestMatcherConfig) NewMatcher() (matcher HttpMatcher, err error) {

	httpMatcher := HttpRequestMatcher{}

	if len(m.HostRegexString) > 0 {
		myHostMatcher, err := NewHostHttpMatcher(m.HostRegexString)
		if err != nil {
			return nil, errors.New("Error validating YAML text: " + err.Error())
		}
		httpMatcher.Matchers = append(httpMatcher.Matchers, myHostMatcher)
	}

	if len(m.VerbRegexString) > 0 {
		myVerbMatcher, err := NewVerbHttpMatcher(m.VerbRegexString)
		if err != nil {
			return nil, errors.New("Error validating YAML text: " + err.Error())
		}
		httpMatcher.Matchers = append(httpMatcher.Matchers, myVerbMatcher)
	}

	if len(m.PathRegexString) > 0 {
		myPathMatcher, err := NewUrlPathHttpMatcher(m.PathRegexString)
		if err != nil {
			return nil, errors.New("Error validating YAML text: " + err.Error())
		}
		httpMatcher.Matchers = append(httpMatcher.Matchers, myPathMatcher)
	}

	for _, param := range m.QueryParams {

		paramMatcher, err := param.NewMatcher()
		if err != nil {
			return nil, errors.New("Error validating YAML text: " + err.Error())
		}

		httpMatcher.Matchers = append(httpMatcher.Matchers, paramMatcher)
	}
	return httpMatcher, err
}

func (m KeyValuesHttpMatcherConfig) NewMatcher() (matcher QueryParamHttpMatcher, err error) {

	myKeyRegex, err := regexp.Compile(m.KeyRegexString)
	if err != nil {
		return QueryParamHttpMatcher{}, errors.New("Error validating YAML text: " + err.Error())
	}
	myValueRegex, err := regexp.Compile(m.ValueRegexString)
	if err != nil {
		return QueryParamHttpMatcher{}, errors.New("Error validating YAML text: " + err.Error())
	}

	return QueryParamHttpMatcher{KeyValuesHttpMatcher{myKeyRegex, myValueRegex}}, err
}
