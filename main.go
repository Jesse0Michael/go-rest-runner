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
	verbose := flag.Int("v", 0, "verbosity level.\n0 (default): quiet\n1: write response details\n2: write response body")
	flag.Parse()

	requestsFile := flag.Arg(0)
	// Prepare rest runner
	if requestsFile == "" {
		usage()
	}
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
	r, _ := json.Marshal(report)
	log.Print(string(r))
}
