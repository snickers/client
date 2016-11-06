package client

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {

	Context("Test Client", func() {
		It("should create a instance of a Client", func() {
			client, err := NewClient("http://localhost.api.snickers/")
			Expect(err).NotTo(HaveOccurred())
			Expect(client.Endpoint).To(Equal("http://localhost.api.snickers/"))
		})
	})
})
