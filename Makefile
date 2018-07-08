.PHONY: build run

build:
	go build -o ./dist/main

run:
	go build -o ./dist/main
	./dist/main
