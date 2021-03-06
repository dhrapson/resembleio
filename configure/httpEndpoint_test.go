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

var _ = Describe("HttpEndpoint", func() {

	var (
		matchers []HttpMatcher
		endpoint HttpEndpoint
	)
	JustBeforeEach(func() {
		endpoint = NamedHttpEndpoint{"myendpoint", matchers}
	})

	Context("when no matchers are provided", func() {
		BeforeEach(func() {
			matchers = []HttpMatcher{}
		})

		It("should return unmatched", func() {
			req, _ := http.NewRequest("GET", "/", nil)
			Expect(endpoint.Match(req)).To(Equal(false))
		})
	})

	Context("when the request is not matched", func() {

		BeforeEach(func() {
			matchers = []HttpMatcher{NegativeHttpMatcher{}}
		})

		It("should return unmatched", func() {
			req, _ := http.NewRequest("GET", "/", nil)
			Expect(endpoint.Match(req)).To(Equal(false))
		})
	})

	Context("when the request is matched", func() {

		BeforeEach(func() {
			matchers = []HttpMatcher{PositiveHttpMatcher{}}
		})

		It("should return unmatched", func() {
			req, _ := http.NewRequest("GET", "/", nil)
			Expect(endpoint.Match(req)).To(Equal(true))
		})
	})

})
