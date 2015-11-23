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
	// "gopkg.in/yaml.v2"
	"io/ioutil"
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

			It("should raise an error", func() {
				err = config.Parse(configData)
				Expect(err).To(HaveOccurred())
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

		Context("when it contains a HTTP type", func() {

			BeforeEach(func() {
				filename = "fixtures/http_resemble.yml"
			})

			It("should return a configuration", func() {
				err = config.Parse(configData)
				Expect(err).NotTo(HaveOccurred())
				Expect(config.TypeName).To(Equal("HTTP"))
			})
		})
	})
})
