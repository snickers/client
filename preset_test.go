package client

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	"net/http"
)

var _ = Describe("Presets", func() {
	var (
		baselinePreset Preset
		rawJSONPreset  string
	)

	BeforeEach(func() {
		rawJSONPreset = `{
		  "name": "baseline_preset",
		  "description": "Baseline Preset",
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
		}`

		json.Unmarshal([]byte(rawJSONPreset), &baselinePreset)

	})

	It("should delete a preset given a preset name", func() {
		server := StartFakeServer(http.StatusOK, rawJSONPreset)
		defer server.Close()
		client, _ := NewClient(server.URL)
		err := client.DeletePreset("baseline_preset")
		Expect(err).NotTo(HaveOccurred())
	})

	It("should return a list of presets", func() {
		expectedPreset := []Preset{
			baselinePreset,
		}
		server := StartFakeServer(http.StatusOK, "["+rawJSONPreset+"]")
		defer server.Close()
		client, _ := NewClient(server.URL)
		respPreset, err := client.GetPresets()
		Expect(respPreset).To(Equal(expectedPreset))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should get a preset details given a preset name", func() {
		server := StartFakeServer(http.StatusOK, rawJSONPreset)
		defer server.Close()
		client, _ := NewClient(server.URL)
		respPreset, err := client.GetPreset("baseline_preset")
		Expect(respPreset).To(Equal(&baselinePreset))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create a preset", func() {
		server := StartFakeServer(http.StatusOK, rawJSONPreset)
		defer server.Close()
		client, _ := NewClient(server.URL)
		respPreset, err := client.CreatePreset(baselinePreset)
		Expect(respPreset).To(Equal(&baselinePreset))
		Expect(err).NotTo(HaveOccurred())
	})

})
