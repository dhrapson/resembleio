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
	"io/ioutil"
	"regexp"
)

var _ = Describe("ResembleConfig", func() {

	var (
		configData []byte
		err        error
		config     ResembleConfig
		filename   string
	)

	JustBeforeEach(func() {
		configData, err = ioutil.ReadFile(filename)
	})

	Describe("reading an invalid config yaml", func() {

		Context("when it contains garbage", func() {

			BeforeEach(func() {
				filename = "fixtures/invalid_config.yml"
			})

			It("should raise a worthy error", func() {
				err = config.Parse(configData)
				Expect(err).To(HaveOccurred())
				matched, matching_err := regexp.MatchString("Error reading YAML text", err.Error())
				Expect(matching_err).NotTo(HaveOccurred())
				Expect(matched).To(BeTrue())
			})
		})

		Context("when it contains no value for type", func() {

			BeforeEach(func() {
				filename = "fixtures/missing_type_config.yml"
			})

			It("should raise an error", func() {
				err = config.Parse(configData)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("reading a valid config yaml", func() {

		Context("when it contains a minimal HTTP type", func() {

			BeforeEach(func() {
				filename = "fixtures/http_resemble.yml"
			})

			It("should return a configuration", func() {
				err = config.Parse(configData)
				Expect(err).NotTo(HaveOccurred())
				Expect(config.TypeName).To(Equal("HTTP"))
			})
		})

		Context("when it contains a partial HTTP type without matchers", func() {

			BeforeEach(func() {
				filename = "fixtures/http_no_matchers_resemble.yml"
			})

			It("should return a configuration", func() {
				err = config.Parse(configData)
				Expect(err).NotTo(HaveOccurred())
				Expect(config.TypeName).To(Equal("HTTP"))
			})
		})

		Context("when it contains a partial HTTP type without query params", func() {

			BeforeEach(func() {
				filename = "fixtures/http_no_query_params_resemble.yml"
			})

			It("should return a configuration", func() {
				err = config.Parse(configData)
				Expect(err).NotTo(HaveOccurred())
				Expect(config.TypeName).To(Equal("HTTP"))
				pathMatcher, _ := NewUrlPathHttpMatcher("/test")
				Expect(config.Matchers[0]).To(Equal(pathMatcher))
				verbMatcher, _ := NewVerbHttpMatcher("GET|POST")
				Expect(config.Matchers[1]).To(Equal(verbMatcher))
				hostMatcher, _ := NewHostHttpMatcher("localhost")
				Expect(config.Matchers[2]).To(Equal(hostMatcher))
			})
		})

		Context("when it contains an HTTP type corrupt query params", func() {

			BeforeEach(func() {
				filename = "fixtures/http_corrupt_params_resemble.yml"
			})

			It("should raise an error", func() {
				err = config.Parse(configData)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when it contains a full HTTP type", func() {

			BeforeEach(func() {
				filename = "fixtures/http_full_resemble.yml"
			})

			It("should return a configuration", func() {
				err = config.Parse(configData)
				Expect(err).NotTo(HaveOccurred())
				Expect(config.TypeName).To(Equal("HTTP"))
				pathMatcher, _ := NewUrlPathHttpMatcher("/test")
				Expect(config.Matchers[0]).To(Equal(pathMatcher))
				verbMatcher, _ := NewVerbHttpMatcher("GET|POST")
				Expect(config.Matchers[1]).To(Equal(verbMatcher))
				hostMatcher, _ := NewHostHttpMatcher("localhost")
				Expect(config.Matchers[2]).To(Equal(hostMatcher))
				paramMatcher1, _ := NewQueryParamHttpMatcher("guid", "[a-zA-Z-0-9-]*")
				Expect(config.Matchers[3]).To(Equal(paramMatcher1))
				paramMatcher2, _ := NewQueryParamHttpMatcher("abc", "123")
				Expect(config.Matchers[4]).To(Equal(paramMatcher2))
			})
		})
	})
})
