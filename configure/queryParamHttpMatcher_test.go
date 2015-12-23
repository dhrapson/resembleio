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
	"net/http"
	"regexp"
)

var _ = Describe("QueryParamHttpMatcher", func() {

	var (
		keyRegex   string
		valueRegex string
		matcher    QueryParamHttpMatcher
		req        *http.Request
		err        error
	)

	JustBeforeEach(func() {
		key, _ := regexp.Compile(keyRegex)
		value, _ := regexp.Compile(valueRegex)
		matcher = QueryParamHttpMatcher{KeyValuesHttpMatcher{key, value}}
	})

	Describe("matching", func() {
		Context("when given an exactly matching regexp", func() {

			BeforeEach(func() {
				keyRegex = `id`
				valueRegex = `^123456$`
				req, _ = http.NewRequest("GET", "/?id=123456", nil)

			})

			It("should not match or throw an error", func() {
				Expect(err).NotTo(HaveOccurred())
				result := matcher.Match(req)
				Expect(result).To(BeTrue())
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})

	Context("when given a non-matching id", func() {

		BeforeEach(func() {
			keyRegex = `id`
			valueRegex = `123456`
			req, _ = http.NewRequest("GET", "/?id=abc", nil)
		})

		It("should return false", func() {
			Expect(err).NotTo(HaveOccurred())
			result := matcher.Match(req)
			Expect(result).To(BeFalse())
			Expect(err).NotTo(HaveOccurred())
		})
	})

})
