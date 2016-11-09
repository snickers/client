package client

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Client", func() {
	It("should create a instance of a Client", func() {
		xclient, err := NewClient("http://localhost.api.snickers/")
		Expect(err).NotTo(HaveOccurred())
		Expect(xclient.Endpoint).To(Equal("http://localhost.api.snickers/"))
	})

	It("should execute do generic method", func() {
		var req *http.Request
		var data []byte
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req = r
			data, _ = ioutil.ReadAll(r.Body)
			w.Write([]byte(`[{"name":"mp4_240p","description":"Encodes video in H264/MP4 @ 240p","container":"mp4","rateControl":"vbr","video":{"width":"426","height":"240","codec":"h264","bitrate":"1000000","gopSize":"90","gopMode":"fixed","profile":"main","profileLevel":"3.1","interlaceMode":"progressive"},"audio":{"codec":"aac","bitrate":"64000"}}]`))
		}))
		defer server.Close()
		var respObj []Preset

		client, _ := NewClient("http://localhost:8000")
		err := client.do("GET", "/presets", []interface{}{}, &respObj)
		Expect(respObj[0].Name).To(Equal("mp4_240p"))
		Expect(err).NotTo(HaveOccurred())

	})
})
