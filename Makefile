.PHONY: build run

build:
	cd src; go build -o ../dist/main

run:
	cd src; go build -o ../dist/main; cd ..
	./dist/main
