all: deps build

deps:
  go get main

build:
  GOOS=linux GOARCH=amd64 go build -o main main
