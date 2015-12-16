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

package configure_test

import (
	. "github.com/dhrapson/resembleio/configure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"regexp"
	"io/ioutil"
)

var _ = Describe("Configure", func() {

	var (
		serviceType ServiceType
		yaml        string
		err         error
	)

	Describe("ConfigureService", func() {
		JustBeforeEach(func() {
			serviceType, err = ConfigureService(yaml)
		})

		Context("when no config yaml is provided", func() {
			BeforeEach(func() {
				yaml = ``
			})

			It("should return an ApiOnlyServiceType", func() {
				Expect(serviceType.Name()).To(Equal("empty service"))
			})
		})

		Context("when a REST config yaml is provided", func() {
			BeforeEach(func() {
				yaml = `---
	type: REST_HTTP`
			})

			It("should return an RestServiceType", func() {
				Expect(serviceType.Name()).To(Equal("REST over HTTP"))
			})
		})

		Context("when a HTTP config yaml is provided", func() {
			BeforeEach(func() {
				yaml = `---
	type: HTTP`
			})

			It("should return an HttpServiceType", func() {
				Expect(serviceType.Name()).To(Equal("HTTP"))
			})
		})
	})

	var _ = Describe("ReadHttpConfiguration", func() {

		var (
			configData string
			err        error
			config     ResembleConfig
			filename   string
		)

		JustBeforeEach(func() {
			configBytes, _ := ioutil.ReadFile(filename)
			configData = string(configBytes)
		})

		Describe("reading an invalid config yaml", func() {

			Context("when it contains garbage", func() {

				BeforeEach(func() {
					filename = "fixtures/invalid_config.yml"
				})

				It("should raise a worthy error", func() {
					_, err = ReadHttpConfiguration(configData)
					Expect(err).To(HaveOccurred())
					matched, matching_err := regexp.MatchString("Error reading YAML text", err.Error())
					Expect(matching_err).NotTo(HaveOccurred())
					Expect(matched).To(BeTrue())
				})
			})

		})

		Describe("reading a valid config yaml", func() {

			Context("when it contains a minimal HTTP type", func() {

				BeforeEach(func() {
					filename = "fixtures/http_resemble.yml"
				})

				It("should return a configuration", func() {
					config, err = ReadHttpConfiguration(configData)
					Expect(err).NotTo(HaveOccurred())
					Expect(config.TypeName).To(Equal("HTTP"))
				})
			})


			Context("when it contains a full HTTP type", func() {

				BeforeEach(func() {
					filename = "fixtures/http_full_resemble.yml"
				})

				It("should return a configuration", func() {
					config, err = ReadHttpConfiguration(configData)
					Expect(err).NotTo(HaveOccurred())
					Expect(config.TypeName).To(Equal("HTTP"))
					httpRequestMatcher0 := config.EndpointConfigs[0].MatcherConfigs[0]
					Expect(httpRequestMatcher0.PathRegexString).To(Equal("/test"))
					Expect(httpRequestMatcher0.VerbRegexString).To(Equal("GET|POST"))
					Expect(httpRequestMatcher0.HostRegexString).To(Equal("localhost"))
					Expect(httpRequestMatcher0.QueryParams[0].KeyRegexString).To(Equal("guid"))
					Expect(httpRequestMatcher0.QueryParams[0].ValueRegexString).To(Equal("[a-zA-Z-0-9-]*"))
					Expect(httpRequestMatcher0.QueryParams[1].KeyRegexString).To(Equal("abc"))
					Expect(httpRequestMatcher0.QueryParams[1].ValueRegexString).To(Equal("123"))
					Expect(httpRequestMatcher0.Headers[0].KeyRegexString).To(Equal("Accept"))
					Expect(httpRequestMatcher0.Headers[0].ValueRegexString).To(Equal("text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"))
					httpRequestMatcher1 := config.EndpointConfigs[0].MatcherConfigs[1]
					Expect(httpRequestMatcher1.PathRegexString).To(Equal("/testagain"))
					Expect(httpRequestMatcher1.VerbRegexString).To(Equal("PUT"))
					Expect(httpRequestMatcher1.HostRegexString).To(Equal("somehost"))
					httpResponder0 := config.EndpointConfigs[0].ResponderConfigs[0]
					Expect(httpResponder0.Name).To(Equal("optional_name"))
					Expect(httpResponder0.Mode).To(Equal("normal"))
				})
			})
		})
	})

})
