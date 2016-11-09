// Package client provides types and methods for interacting with the
// Snickers API.
//
package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client is the basic type for interacting with the API. It provides methods
// matching the available actions in the API.
type Client struct {
	Endpoint string
}

// NewClient creates a instance of the client type.
func NewClient(endpoint string) (*Client, error) {
	return &Client{Endpoint: endpoint}, nil
}

// APIError represents an error returned by the Snickers API.
//
type APIError struct {
	Status int    `json:"status,omitempty"`
	Errors string `json:"errors,omitempty"`
}

// Error converts the whole interlying information to a representative string.
//
// It encodes the list of errors in JSON format.
func (apiErr *APIError) Error() string {
	data, _ := json.Marshal(apiErr)
	return fmt.Sprintf("Error returned by the Snickers API: %s", data)
}

func (c *Client) do(method string, path string, body interface{}, out interface{}) error {

	jsonRequest, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, c.Endpoint+path, strings.NewReader(string(jsonRequest)))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return &APIError{
			Status: resp.StatusCode,
			Errors: string(respData),
		}
	}

	if out != nil && len(respData) > 1 {
		return json.Unmarshal(respData, out)
	}

	return nil
}
