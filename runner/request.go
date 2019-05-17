package runner

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

// Request is a structure containing a web request that will be made
type Request struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Body    RequestBody       `json:"body,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Times   int               `json:"times,omitempty"`
}

// RequestBody allows control over the Request Body's encoding
type RequestBody []byte

// UnmarshalJSON is a custom implementation for JSON Unmarshalling for the RequestBody
// Unmarshalling will first check if the data is a local filepath that can be read
// Else it will check if the data is stringified JSON and un-stringify the data to use
// or Else it will just use the []byte
func (response *RequestBody) UnmarshalJSON(data []byte) error {
	unmarshaled := []byte{}
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		// The data is a []byte, so use it
		unmarshaled = data
	}

	if s, err := strconv.Unquote(string(unmarshaled)); err == nil {
		absPath, _ := filepath.Abs(s)
		if _, err := os.Stat(absPath); err == nil {
			// The data is a path that exists, therefore we will read the file
			if file, err := ioutil.ReadFile(absPath); err == nil {
				*response = file
				return nil
			}
		}
		// The data is stringified JSON, therefore we eill use the unquoted JSON
		*response = []byte(s)
		return nil
	}

	*response = unmarshaled
	return nil
}
