package client

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
)

var _ = Describe("Client", func() {
	It("should create a instance of a Client", func() {
		xclient, err := NewClient("http://localhost.api.snickers/")
		Expect(err).NotTo(HaveOccurred())
		Expect(xclient.Endpoint).To(Equal("http://localhost.api.snickers/"))
	})

	It("should execute do generic method", func() {
		server := StartFakeServer(http.StatusOK, `[{
				  "name": "mp4_240p",
				  "description": "Test Preset",
				  "container": "mp4",
				  "rateControl": "vbr",
				  "video": {
				      "height": "720",
				      "width": "1280",
				      "codec": "h264",
				      "bitrate": "10000"
				   },
				  "audio": {
				      "codec": "aac",
				      "bitrate": "64000"
				  }
			}]`)
		defer server.Close()
		client, _ := NewClient(server.URL)

		var respObj []Preset
		err := client.do("GET", "/presets", []interface{}{}, &respObj)
		Expect(respObj[0].Name).To(Equal("mp4_240p"))
		Expect(err).NotTo(HaveOccurred())

	})
})
