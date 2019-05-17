package runner

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"time"
)

// Client is the struct that controls the actions of the rest runner
type Client struct {
	Requests   []Request
	httpClient http.Client

	Verbose int
}

// NewClient creates a new rest runner client
func NewClient(requests []Request) *Client {
	return &Client{
		Requests:   requests,
		httpClient: *http.DefaultClient,
	}
}

// Run executes every request, records data, and returns a report
func (c *Client) Run() ([]GroupReport, error) {
	report := []GroupReport{}
	for _, r := range c.Requests {
		call := fmt.Sprintf("%s:%s", r.Method, r.URL)
		responses := Responses{}
		for i := 0; i < int(math.Max(1, float64(r.Times))); i++ {
			start := time.Now()
			req, err := http.NewRequest(r.Method, r.URL, bytes.NewReader(r.Body))
			if err != nil {
				return nil, err
			}

			for key, val := range r.Headers {
				req.Header.Add(key, val)
			}
			resp, err := c.httpClient.Do(req)
			if err != nil {
				return nil, err
			}

			duration := time.Since(start).Seconds()
			responses = append(responses, Response{
				Call:       call,
				StatusCode: resp.StatusCode,
				Duration:   duration,
			})

			if c.Verbose >= 1 {
				log.Printf("%s %s %d %fs\n", r.Method, r.URL, resp.StatusCode, duration)
			}

			if c.Verbose >= 2 {
				body, _ := ioutil.ReadAll(resp.Body)
				log.Println(string(body))
			}
		}

		report = append(report, responses.GroupReport())
	}

	return report, nil
}
