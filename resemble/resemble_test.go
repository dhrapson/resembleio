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

package main_test

import (
	. "github.com/dhrapson/resembleio/resemble"
	"os/exec"
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func copyFile(srcFolder, destFolder string) {
	cpCmd := exec.Command("cp", "-f", srcFolder, destFolder)
	err := cpCmd.Run()
	if err != nil {
		panic(err)
	}
}

func deleteFile(filePath string) {
	cpCmd := exec.Command("rm", "-f", filePath)
	err := cpCmd.Run()
	if err != nil {
		panic(err)
	}
}

var _ = Describe("Resemble", func() {

	Describe("loading config data from file", func() {

		var (
			configString string
			cmdLineArgs  []string
			err          error
		)

		JustBeforeEach(func() {
			configString, err = GetConfigData(cmdLineArgs)
		})

		Context("when no config yaml is provided", func() {

			BeforeEach(func() {
				cmdLineArgs = []string{}
			})

			Context("and no config yaml is found in the default location", func() {
				It("should return empty string", func() {
					Expect(err).NotTo(HaveOccurred())
					Expect(configString).To(Equal(""))
				})
			})

			Context("and a valid REST config yaml is found in the default location", func() {
				BeforeEach(func() {
					copyFile("fixtures/rest_resemble.yml", "resemble.yml")
				})

				AfterEach(func() {
					deleteFile("resemble.yml")
				})

				It("should read the file contents", func() {
					Expect(err).NotTo(HaveOccurred())
					matched, matching_err := regexp.MatchString("REST_HTTP", configString)
					Expect(matching_err).NotTo(HaveOccurred())
					Expect(matched).To(BeTrue())
				})
			})
		})

		Context("when a config yaml path is provided", func() {

			Context("and no file is found in the provided location", func() {

				BeforeEach(func() {
					cmdLineArgs = []string{"wont_ever_exist.yml"}
				})

				It("should error", func() {
					Expect(err).To(HaveOccurred())
				})
			})

			Context("and a file is found in the provided location", func() {

				BeforeEach(func() {
					cmdLineArgs = []string{"fixtures/rest_resemble.yml"}
				})

				It("should read the file contents", func() {
					Expect(err).NotTo(HaveOccurred())
					matched, matching_err := regexp.MatchString("REST_HTTP", configString)
					Expect(matching_err).NotTo(HaveOccurred())
					Expect(matched).To(BeTrue())
				})
			})
		})
	})
})
