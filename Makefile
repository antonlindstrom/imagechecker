.DEFAULT_GOAL := build

LINT_GOMETALINTER_OK := $(shell which gometalinter 2> /dev/null)

build: bin/imagechecker

bin/imagechecker: cmd/imagechecker/main.go
	mkdir -p bin/
	go build -o bin/imagechecker cmd/imagechecker/main.go

.PHONY: clean
clean:
	-rm -r bin/

.PHONY: check
check:
	go test -race -cover -v ./...

.PHONY: lint
lint:
ifndef LINT_GOMETALINTER_OK
	go get github.com/alecthomas/gometalinter
	gometalinter --install
endif
	gometalinter --disable=gotype ./... --fast --deadline=10s
