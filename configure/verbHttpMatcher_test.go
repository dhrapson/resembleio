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

var _ = Describe("VerbHttpMatcher", func() {

	var (
		regex   string
		matcher VerbHttpMatcher
		err     error
	)

	Describe("when using a valid regex", func() {

		JustBeforeEach(func() {
			matcher, err = NewVerbHttpMatcher(regex)
		})

		Context("when given an exactly matching regexp", func() {

			BeforeEach(func() {
				regex = `GET`
			})

			It("should return true", func() {
				Expect(err).NotTo(HaveOccurred())
				result := matcher.MatchVerb("GET")
				Expect(result).To(BeTrue())
			})
		})

		Context("when given a loosely matching regexp", func() {

			BeforeEach(func() {
				regex = `GET|POST|DELETE`
			})

			It("should return true", func() {
				Expect(err).NotTo(HaveOccurred())
				result := matcher.MatchVerb("POST")
				Expect(result).To(BeTrue())
			})
		})

		Context("when given a non-matching verb", func() {

			BeforeEach(func() {
				regex = `GET|DELETE`
			})

			It("should return false", func() {
				Expect(err).NotTo(HaveOccurred())
				result := matcher.MatchVerb("POST")
				Expect(result).To(BeFalse())
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})

	Describe("when using an invalid regex", func() {
		BeforeEach(func() {
			regex = `^abc++$`
		})

		It("should raise an error on creating the matcher", func() {
			matcher, err = NewVerbHttpMatcher(regex)
			Expect(err).To(HaveOccurred())
		})
	})

})
