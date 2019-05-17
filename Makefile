GO ?= go
COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
	COVERAGEDIR=$(CIRCLE_ARTIFACTS)/coverage
endif

LDFLAGS = -ldflags '-X main.gitSHA=$(shell git rev-parse HEAD)'

all: build test cover
build:
	if [ ! -d bin ]; then mkdir bin; fi
	$(GO) build $(LDFLAGS) -v -o bin/go-rest-runner
test:
	if [ ! -d $(COVERAGEDIR) ]; then mkdir $(COVERAGEDIR); fi
	$(GO) test -v ./runner -cover -coverprofile=$(COVERAGEDIR)/runner.coverprofile
cover:
	if [ ! -d $(COVERAGEDIR) ]; then mkdir $(COVERAGEDIR); fi
	$(GO) tool cover -html=$(COVERAGEDIR)/runner.coverprofile -o $(COVERAGEDIR)/runner.html
coveralls:
	if [ ! -d $(COVERAGEDIR) ]; then mkdir $(COVERAGEDIR); fi
	gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile
	goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile  -service=circle-ci -repotoken=$(COVERALLS_TOKEN)
assert-no-diff:
	test -z "$(shell git status --porcelain)"
clean:
	$(GO) clean
	rm -f bin/go-rest-runner
	rm -rf coverage/
