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

package serve_test

import (
	. "github.com/dhrapson/resembleio/serve"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockServiceType struct {
	served bool
}

func (m *MockServiceType) Configure() {}

func (m *MockServiceType) Name() string {
	return "mock"
}
func (m *MockServiceType) Serve() {
	m.served = true
}

var _ = Describe("Serve", func() {

	var (
		serviceType *MockServiceType
	)

	BeforeEach(func() {
		serviceType = new(MockServiceType)
	})

	It("calls serve on the ServiceType", func() {
		Serve(serviceType)
		serviceType.Configure()
		Expect(serviceType.served).To(BeTrue())
	})
})
