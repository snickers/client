package client

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	"net/http"
)

var _ = Describe("Jobs", func() {
	var (
		jobInput      JobInput
		job           Job
		jobJSONOutput string
	)
	BeforeEach(func() {
		jobJSONOutput = `{
			"id":"5iNcv6pMWYVCZhRL",
			"source":"http://flv.io/KailuaBeach.mp4",
			"destination":"http://AKIAII3C6HP:4wooQI51NXIN5ELJNTq@snickers-media.s3.amazonaws.com/outputs/",
			"preset":{
				"name":"mp4_240p",
				"description":"Encodes video in H264/MP4 @ 240p",
				"container":"mp4",
				"rateControl":"vbr",
				"video":{
					"width":"426",
					"height":"240",
					"codec":"h264",
					"bitrate":"1000000",
					"gopSize":"90",
					"gopMode":"fixed",
					"profile":"main",
					"profileLevel":"3.1",
					"interlaceMode":"progressive"
				},
				"audio":{
					"codec":"aac",
					"bitrate":"64000"
				}
			},
			"status":"created",
			"progress":""
		}`

		json.Unmarshal([]byte(jobJSONOutput), &job)

		jobInput = JobInput{
			Source:      "http://flv.io/KailuaBeach.mp4",
			Destination: "http://AKIAII3C6HP:4wooQI51NXIN5ELJNTq@snickers-media.s3.amazonaws.com/outputs/",
			PresetName:  "mp4_240p",
		}
	})

	It("should return a list of jobs", func() {
		expectedJobs := []Job{
			job,
		}
		server := StartFakeServer(http.StatusOK, "["+jobJSONOutput+"]")
		defer server.Close()
		client, _ := NewClient(server.URL)
		respJob, err := client.GetJobs()
		Expect(respJob).To(Equal(expectedJobs))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should return a given job details", func() {
		server := StartFakeServer(http.StatusOK, jobJSONOutput)
		defer server.Close()
		client, _ := NewClient(server.URL)
		respJob, err := client.GetJob("5iNcv6pMWYVCZhRL")
		Expect(respJob).To(Equal(&job))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should start a job given a job id", func() {
		server := StartFakeServer(http.StatusOK, "")
		defer server.Close()
		client, _ := NewClient(server.URL)
		_, err := client.StartJob("5iNcv6pMWYVCZhRL")
		Expect(err).NotTo(HaveOccurred())
	})

	It("should create a job", func() {
		server := StartFakeServer(http.StatusOK, jobJSONOutput)
		defer server.Close()
		client, _ := NewClient(server.URL)
		respJob, err := client.CreateJob(jobInput)
		Expect(respJob).To(Equal(&job))
		Expect(err).NotTo(HaveOccurred())
	})

})
