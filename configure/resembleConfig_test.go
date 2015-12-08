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
)

var _ = Describe("ResembleConfig", func() {

	var (
	  verbRegex   string
		hostRegex   string
		pathRegex   string
		queryParams [][]string
		matcher HttpMatcher
		matcherConfig HttpRequestMatcherConfig
		err     error
	)

	Describe("creating a new HttpRequestMatcher", func() {

		JustBeforeEach(func() {
			queryParamConfigs := []KeyValuesHttpMatcherConfig{}
			for _, queryParam := range queryParams {
				queryParamConfigs = 	append(queryParamConfigs, KeyValuesHttpMatcherConfig{queryParam[0], queryParam[1]})
			}
			matcherConfig = HttpRequestMatcherConfig{"name", verbRegex, hostRegex, pathRegex, queryParamConfigs}
		})

		Context("when no regexs are provided", func() {

			It("should not throw an error", func() {
				matcher, err = matcherConfig.NewMatcher()
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when all valid regexs are provided", func() {

			BeforeEach(func() {
				verbRegex = "abc"
				hostRegex = "123"
				pathRegex   = "def"
				queryParams = [][]string{{"name1","value1"}, {"name2","value2"}}
			})

			It("should not throw an error", func() {
				matcher, err = matcherConfig.NewMatcher()
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when an invalid regex is provided", func() {

			BeforeEach(func() {
				verbRegex = "abc"
				hostRegex = `^abc++$`
				pathRegex   = "def"
				queryParams = [][]string{{"name1","value1"}, {"name2","value2"}}
			})

			It("should throw an error", func() {
				matcher, err = matcherConfig.NewMatcher()
				Expect(err).To(HaveOccurred())
			})
		})

	})

})

