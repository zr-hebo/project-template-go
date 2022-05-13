BINARY_NAME=project-template
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

build:
	go env -w GO111MODULE="on"
	go env -w GOPROXY="https://goproxy.cn,direct"
	$(GOBUILD) -o bin/$(BINARY_NAME) -v cmd/*.go

test:
	go test github.com/zr-hebo/project-template-go/internal
	cd scripts/test && python3 test.py
