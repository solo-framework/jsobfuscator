# .DEFAULT_GOAL := help
.PHONY: *

LOCAL_BUILD_DIR := ./build
TARGETOS := linux
TARGETARCH := amd64
# linux/amd64

build:
	@rm -rf ./build/ \
	&& mkdir ./build \
	&& go mod download \
	&& go mod tidy \
	&& CGO_ENABLED=1 GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) \
		go build -trimpath  -ldflags "-s -w" -o ./build/jsobfs ./app

