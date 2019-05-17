package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jesse0michael/go-rest-runner/runner"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage of go-rest-runner:")
	fmt.Fprintln(os.Stderr, "go-rest-runner [options..] path/to/requests_file.json")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	verbose := flag.Int("v", 0, "verbosity level. [0: quiet | 1: write response details | 2: write response body] (default 0)")
	format := flag.String("f", "text", "report format. [text | json]")
	flag.Parse()

	// Validate arguments
	requestsFile := flag.Arg(0)
	if requestsFile == "" {
		usage()
	}
	if format != nil && (*format != "text" && *format != "json") {
		usage()
	}

	// Prepare rest runner
	file, err := os.Open(requestsFile)
	if err != nil {
		log.Fatalf("failed to open: %s error: %s", requestsFile, err.Error())
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read: %s error: %s", requestsFile, err.Error())
	}
	var requests []runner.Request
	if err := json.Unmarshal(b, &requests); err != nil {
		log.Fatalf("failed to unmarshal: %s error: %s", requestsFile, err.Error())
	}
	client := runner.NewClient(requests)
	client.Verbose = *verbose

	// Run rest runner
	report, err := client.Run()
	if err != nil {
		log.Fatalf("failed to run the rest runner error: %s", err.Error())
	}

	// Report rest runner
	switch *format {
	case "json":
		r, _ := json.Marshal(report)
		log.Println(string(r))
	case "text":
		fmt.Println("Go Rest Runner")
		fmt.Println("--------------------------------------------------")
		for _, r := range report {
			fmt.Printf("%s times: %d average duration: %f status: %s\n", r.Call, r.Times, r.AvgDuration, r.StatusCodes.String())
		}
	}
}
