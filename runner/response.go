package runner

import (
	"fmt"
)

// Response contains data on a call made by the rest runner
type Response struct {
	Call       string  `json:"call"`
	StatusCode int     `json:"status_code"`
	Duration   float64 `json:"duration"`
}

// Responses define a collection of rest runner responses
type Responses []Response

// GroupReport turns a slice of rest runner responses of the same call and analizes them as a group
func (rs Responses) GroupReport() GroupReport {
	codes := map[int]int{}
	var sumDuration float64

	for _, r := range rs {
		sumDuration = sumDuration + r.Duration
		count, _ := codes[r.StatusCode]
		codes[r.StatusCode] = count + 1
	}

	groupCodes := []GroupStatusCode{}
	for code, count := range codes {
		groupCodes = append(groupCodes, GroupStatusCode{
			Code:  code,
			Times: count,
		})
	}

	group := GroupReport{
		Times:       len(rs),
		StatusCodes: groupCodes,
	}

	if len(rs) > 0 {
		group.Call = rs[0].Call
		group.AvgDuration = sumDuration / float64(len(rs))
	}

	return group
}

// GroupReport is a collection of analized responses for the same call
type GroupReport struct {
	Call        string      `json:"call"`
	Times       int         `json:"times"`
	AvgDuration float64     `json:"avg_duration"`
	StatusCodes StatusCodes `json:"status_codes"`
}

// StatusCodes define a collection of status codes for a group of calls
type StatusCodes []GroupStatusCode

func (sc StatusCodes) String() string {
	var s string
	for _, status := range sc {
		s = s + status.String() + " "
	}
	return s
}

// GroupStatusCode is the analized status codes for a response group
type GroupStatusCode struct {
	Code  int `json:"code"`
	Times int `json:"times"`
}

func (sc GroupStatusCode) String() string {
	s := fmt.Sprintf("%d", sc.Code)
	if sc.Times > 1 {
		s = s + fmt.Sprintf(" x%d ", sc.Times)
	}
	return s
}
