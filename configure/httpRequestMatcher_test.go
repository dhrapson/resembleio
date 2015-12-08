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
)

var _ = Describe("HttpRequestMatcher", func() {

	var _ = Describe("#Match", func() {

		var (
		  matchers []HttpMatcher
			matcher HttpRequestMatcher
			req *http.Request
			err     error
		)

		JustBeforeEach(func() {
			req, _ = http.NewRequest("GET", "http://localhost/", nil)
			matcher = HttpRequestMatcher{"name", matchers}
		})

		BeforeEach(func() {
			matchers = []HttpMatcher{}
		})

		Context("when there are no child matchers", func() {

			It("should not match or throw an error", func() {
				Expect(err).NotTo(HaveOccurred())
				result := matcher.Match(req)
				Expect(result).To(BeFalse())
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when there is one positive matcher", func() {

			BeforeEach(func() {
				matchers = []HttpMatcher{PositiveHttpMatcher{}}
			})

			It("should match and not throw an error", func() {
				Expect(err).NotTo(HaveOccurred())
				result := matcher.Match(req)
				Expect(result).To(BeTrue())
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the first matcher is negative", func() {

			BeforeEach(func() {
				matchers = []HttpMatcher{NegativeHttpMatcher{}, PositiveHttpMatcher{}}
			})

			It("should not match and not throw an error", func() {
				Expect(err).NotTo(HaveOccurred())
				result := matcher.Match(req)
				Expect(result).To(BeFalse())
				Expect(err).NotTo(HaveOccurred())
			})

		})
	})
})

