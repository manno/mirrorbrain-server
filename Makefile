export GOPATH := $(shell pwd )
export GOBIN := $(shell pwd )/bin

all:
	cd src/app && go get
	go build -a app
