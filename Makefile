.PHONY: build run install


build:
	go build -o twn .

run:
	go run .

install: build
	cp twn $(shell go env GOPATH)/bin/
