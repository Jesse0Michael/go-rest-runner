# GO REST RUNNER

[![CircleCI](https://circleci.com/gh/Jesse0Michael/go-rest-runner.svg?style=svg)](https://circleci.com/gh/Jesse0Michael/go-rest-runner) [![Coverage Status](https://coveralls.io/repos/github/Jesse0Michael/go-rest-runner/badge.svg)](https://coveralls.io/github/Jesse0Michael/go-rest-runner)

Go Rest Runner is a small GO application intended to make a series of web requests and then report on the results of the requests that were made. The requests that it runs are parsed from a JSON whose filepath is provided as a CLI argument.

## Example file

```json
[
  {
    "method": "POST",
    "url": "https://your.target.com/",
    "body": "{\"test\":\"best\"}"
  }
]
```

## Run, Go, Run

Please have [GO](https://golang.org/) installed on your machine

`go get -u github.com/jesse0michael/go-rest-runner`

```bash
Usage of go-rest-runner:
go-rest-runner [options..] path/to/requests_file.json
  -v int
        verbosity level.
        0 (default): quiet
        1: write response details
        2: write response body
```

`go-rest-runner example.json`
