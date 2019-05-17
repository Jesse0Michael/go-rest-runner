package runner

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"time"
)

// Client is the struct that controls the actions of the rest runner
type Client struct {
	Requests   []Request
	httpClient http.Client
}

// NewClient creates a new rest runner client
func NewClient(requests []Request) *Client {
	return &Client{
		Requests:   requests,
		httpClient: *http.DefaultClient,
	}
}

// Run executes every request, records data, and returns a report
func (c *Client) Run() (*Report, error) {
	report := Report{GroupReports: map[string]GroupReport{}}
	runStart := time.Now()
	for _, r := range c.Requests {
		call := fmt.Sprintf("%s:%s", r.Method, r.URL)
		responses := Responses{}
		req, err := http.NewRequest(r.Method, r.URL, bytes.NewReader(r.Body))
		if err != nil {
			return nil, err
		}

		for i := 0; i < int(math.Max(1, float64(r.Times))); i++ {
			start := time.Now()
			resp, err := c.httpClient.Do(req)
			if err != nil {
				return nil, err
			}
			responses = append(responses, Response{
				Call:       call,
				StatusCode: resp.StatusCode,
				Duration:   time.Since(start).Seconds(),
			})
		}

		report.GroupReports[call] = responses.GroupReport()
	}

	report.TotalDuration = time.Since(runStart).Seconds()
	return &report, nil
}
