package client

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Presets", func() {
	var (
		client *Client
	)

	BeforeEach(func() {
		client, _ = NewClient("http://localhost:8000")
	})

	It("should create a preset", func() {
		var req *http.Request
		var data []byte
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req = r
			data, _ = ioutil.ReadAll(r.Body)
			w.Write([]byte(`{
				  "name": "test_preset",
				  "description": "Test Preset",
				  "container": "mp4",
				  "rateControl": "vbr",
				  "video": {
				      "height": "720",
				      "width": "1280",
				      "codec": "h264",
				      "bitrate": "10000",
				   },
				  "audio": {
				      "codec": "aac",
				      "bitrate": "64000"
				  }
			}`))
		}))

		defer server.Close()

		preset := Preset{
			Name:        "test_preset",
			Description: "Test Preset",
			Container:   "mp4",
			RateControl: "vbr",
			Video: VideoPreset{
				Height:  "720",
				Width:   "1280",
				Codec:   "h264",
				Bitrate: "10000",
			},
			Audio: AudioPreset{
				Codec:   "aac",
				Bitrate: "64000",
			},
		}
		respPreset, err := client.CreatePreset(preset)
		Expect(respPreset).To(Equal(&preset))
		Expect(err).NotTo(HaveOccurred())
	})

})
