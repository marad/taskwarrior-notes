.PHONY: build run install


build:
	go build -o twn .

run:
	go run .

install: build
	cp twn $(shell go env GOPATH)/bin/

install_hook:
	mkdir -p ~/.task/hooks/
	cp hooks/on-modify-sync-note.sh ~/.task/hooks/
