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

	It("should fail to delete a preset", func() {
		apiError := &APIError{
			Status: http.StatusBadRequest,
			Errors: "Failed to delete a preset",
		}
		server := StartFakeServer(http.StatusBadRequest, "Failed to delete a preset")
		defer server.Close()
		client, _ := NewClient(server.URL)
		err := client.DeletePreset("error_preset")
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(apiError))
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

	It("should fail to retreive a list of presets", func() {
		apiError := &APIError{
			Status: http.StatusBadRequest,
			Errors: "Failed to retrieve a list presets",
		}
		server := StartFakeServer(http.StatusBadRequest, "Failed to retrieve a list presets")
		defer server.Close()
		client, _ := NewClient(server.URL)
		_, err := client.GetPresets()
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(apiError))
	})

	It("should get a preset details given a preset name", func() {
		server := StartFakeServer(http.StatusOK, rawJSONPreset)
		defer server.Close()
		client, _ := NewClient(server.URL)
		respPreset, err := client.GetPreset("baseline_preset")
		Expect(respPreset).To(Equal(&baselinePreset))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should fail to retreive a given preset details", func() {
		apiError := &APIError{
			Status: http.StatusBadRequest,
			Errors: "Failed to retrieve preset details",
		}
		server := StartFakeServer(http.StatusBadRequest, "Failed to retrieve preset details")
		defer server.Close()
		client, _ := NewClient(server.URL)
		_, err := client.GetPreset("error_preset")
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(apiError))
	})

	It("should create a preset", func() {
		server := StartFakeServer(http.StatusOK, rawJSONPreset)
		defer server.Close()
		client, _ := NewClient(server.URL)
		respPreset, err := client.CreatePreset(baselinePreset)
		Expect(respPreset).To(Equal(&baselinePreset))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should fail to create a preset", func() {
		apiError := &APIError{
			Status: http.StatusBadRequest,
			Errors: "Failed to create a preset",
		}
		server := StartFakeServer(http.StatusBadRequest, "Failed to create a preset")
		defer server.Close()
		client, _ := NewClient(server.URL)
		_, err := client.CreatePreset(baselinePreset)
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(apiError))
	})

})
