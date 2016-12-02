OS      := $(shell uname -a | cut -f 1 -d ' ' | tr [:upper:] [:lower:])
ARCH    := $(shell uname -m)
TAG     := $(shell git describe master --tags)
TIMESTAMP := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

all: build

build:
	@go install
	@cd cmd/tlsping && go build -ldflags="-X main.appBuildTime=$(TIMESTAMP) -X main.appVersion=$(TAG)"

release: build
	@echo "Packaging tlsping ${TAG} for ${OS}"
	@cd cmd/tlsping && tar -czf tlsping-${TAG}-${OS}-${ARCH}.tar.gz tlsping

clean:
	@rm -f cmd/tlsping/tlsping-*.tar.gz cmd/tlsping/tlsping

buildall: clean build
