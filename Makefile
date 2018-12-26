.PHONY: default deps clean fmt pretest lint lint-test list vet build test all
SHELL := /bin/bash
BINARY=insights

VERSION=0.1.0
BUILD_TIME=`date +%FT%T%z`

BRANCH=`git rev-parse --abbrev-ref HEAD`
COMMIT=`git rev-parse --short HEAD`

LDFLAGS="-X ${BINARY}.version=${VERSION} -X ${BINARY}.buildtime=${BUILD_TIME} -X ${BINARY}.branch=${BRANCH} -X ${BINARY}.commit=${COMMIT}"

GOX := $(shell which gox 2>/dev/null)
DEP := $(shell which dep 2>/dev/null)
PACKR := $(shell which packr 2>/dev/null)

OSES := "linux windows darwin"
ARCHS := "amd64 386"

default: build

dep:
ifdef DEP
	@dep ensure
else
	@echo 'Dep is not installed'
	@echo 'Installing Dep...'
	@go get -u github.com/golang/dep/cmd/dep
	@dep ensure
endif

pack:
ifdef PACK
	@packr
else
	@echo 'Packr is not installed'
	@echo 'Installing Packr...'
	@go get -u github.com/gobuffalo/packr/packr
	@packr
endif

clean:
	@packr clean
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

pretest:
	@gofmt -d $$(find . -type f -name '*.go' -not -path "./vendor/*") 2>&1 | read; [ $$? == 1 ]

vet:
	@go vet $(go list -f '{{ .ImportPath }}' ./... | grep -v vendor/)

test: pretest vet lint dep
	@go test -v $$(go list -f '{{ .ImportPath }}' ./... | grep -v vendor/) -p=1

build: clean test
	@packr && go build -x -ldflags ${LDFLAGS} -o bin/${BINARY} github.com/umayr/${BINARY}/cmd/${BINARY}

fmt:
	@gofmt -w $$(find . -type f -name '*.go' -not -path "./vendor/*")

lint:
	@go get -v github.com/golang/lint/golint
	@golint ./... | grep -v vendor/ | true

list:
	@go list -f '{{ .ImportPath }}' ./... | grep -v vendor/

cross:
ifdef GOX
	@gox -ldflags=${LDFLAGS} -output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}" -os=${OSES} -arch=${ARCHS} github.com/umayr/${BINARY}/cmd/${BINARY}
else
	@echo 'Gox is not installed'
	@echo 'Installing Gox...'
	@go get -u github.com/mitchellh/gox
	@gox -ldflags=${LDFLAGS} -output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}" -os=${OSES} -arch=${ARCHS} github.com/umayr/${BINARY}/cmd/${BINARY}
endif