.PHONY: build doc fmt lint run test

default: build

build:
	go build -v -o ./bin/lojjebot ./src/

doc:
	godoc -http=:6060 -index

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./src/...

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint ./src

run: build
	./bin/lojjebot

test:
	go test ./src/...
