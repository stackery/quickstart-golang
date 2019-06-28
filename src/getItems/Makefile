.PHONY: deps build

all: deps build

deps:
	go get

build:
	GOOS=linux GOARCH=amd64 go build -o main
