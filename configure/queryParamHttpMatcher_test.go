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

var _ = Describe("QueryParamHttpMatcher", func() {

	var (
		keyRegex    string
		valueRegex  string
		queryParams map[string][]string
		matcher     KeyValuesHttpMatcher
		err         error
	)

	Describe("when using a valid regex", func() {

		JustBeforeEach(func() {
			matcher, err = NewKeyValuesHttpMatcher(keyRegex, valueRegex)
		})

		Context("when given an exactly matching regexp", func() {

			BeforeEach(func() {
				keyRegex = `id`
				valueRegex = `^123456$`
				queryParams = map[string][]string{
					"id": []string{"abc-here", "123456"},
				}
			})

			It("should return true", func() {
				Expect(err).NotTo(HaveOccurred())
				result := matcher.MatchKeyValues(queryParams)
				Expect(result).To(BeTrue())
			})
		})

		Context("when given a loosely matching regexp", func() {

			BeforeEach(func() {
				keyRegex = `id`
				valueRegex = `[0-9A-Za-z-]*`
				queryParams = map[string][]string{
					"id": []string{"123456-abcd"},
				}
			})

			It("should return true", func() {
				Expect(err).NotTo(HaveOccurred())
				result := matcher.MatchKeyValues(queryParams)
				Expect(result).To(BeTrue())
			})
		})

		Context("when there are bunch of parameters", func() {

			BeforeEach(func() {
				keyRegex = `id`
				valueRegex = `[0-9A-Za-z-]*`
				queryParams = map[string][]string{
					"abc": []string{"notaguid+"},
					"def": []string{"alsonotaguid_"},
					"id":  []string{"", "123456-abcd"},
				}
			})

			It("should return true", func() {
				Expect(err).NotTo(HaveOccurred())
				result := matcher.MatchKeyValues(queryParams)
				Expect(result).To(BeTrue())
			})
		})

		Context("when given a non-matching id", func() {

			BeforeEach(func() {
				keyRegex = `id`
				valueRegex = `123456`
				queryParams = map[string][]string{
					"id": []string{"abcd"},
				}
			})

			It("should return false", func() {
				Expect(err).NotTo(HaveOccurred())
				result := matcher.MatchKeyValues(queryParams)
				Expect(result).To(BeFalse())
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})

	Describe("when using invalid regexes", func() {

		Context("when the keyRegex is invalid", func() {

			BeforeEach(func() {
				keyRegex = `^abc++$`
				valueRegex = `123456`
			})

			It("should raise an error on creating the matcher", func() {
				matcher, err = NewKeyValuesHttpMatcher(keyRegex, valueRegex)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when the valueRegex is invalid", func() {

			BeforeEach(func() {
				keyRegex = `123456`
				valueRegex = `^abc++$`
			})

			It("should raise an error on creating the matcher", func() {
				matcher, err = NewKeyValuesHttpMatcher(keyRegex, valueRegex)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
