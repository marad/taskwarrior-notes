.PHONY: build run install


build:
	go build -o twn .

run:
	go run .

install: build
	go install twn