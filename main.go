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
	fmt.Fprintf(os.Stderr, "usage: go-rest-runner [path/to/requests_file.json]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	requestsFile := flag.Arg(0)

	// Prepare rest runner
	if requestsFile == "" {
		log.Fatalf("path to requests file is required.")
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

	report, err := client.Run()
	if err != nil {
		log.Fatalf("failed to run the rest runner error: %s", err.Error())
	}
	r, _ := json.Marshal(report)
	log.Print(string(r))
}
