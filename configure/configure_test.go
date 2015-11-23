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
	. "github.com/dhrapson/resemble/configure"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Configure", func() {

	var (
		serviceType ServiceType
		yaml        string
		err         error
	)

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
