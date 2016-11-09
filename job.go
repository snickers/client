package client

// JobStatus represents the status of a job
type JobStatus string

// Job is the set of parameters of a given job
type Job struct {
	ID               string    `json:"id"`
	Source           string    `json:"source"`
	Destination      string    `json:"destination"`
	Preset           Preset    `json:"preset"`
	Status           JobStatus `json:"status"`
	Details          string    `json:"progress"`
	LocalSource      string    `json:"-"`
	LocalDestination string    `json:"-"`
}

// JobInput stores the information passed from the
// user when creating a job.
type JobInput struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	PresetName  string `json:"preset"`
}

// GetJobs returns a list of jobs
func (c *Client) GetJobs() ([]Job, error) {
	var result []Job
	err := c.do("GET", "/jobs", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetJob returns metadata on a single job
func (c *Client) GetJob(jobID string) (*Job, error) {
	var result *Job
	err := c.do("GET", "/jobs/"+jobID, nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateJob sends a single job and send it for processing
func (c *Client) CreateJob(jobInput JobInput) (*Job, error) {
	var result *Job
	err := c.do("POST", "/jobs", jobInput, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StartJob start a job given a job id
func (c *Client) StartJob(jobID string) (*Job, error) {
	var result *Job
	err := c.do("POST", "/jobs/"+jobID+"/start", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
