.PHONY: build run

build:
	go build -o ./dist/main

run:
	go build -o ./dist/main
	./dist/main -programPath=_resources/Phase1/Prg/4_Colors.prg

debug:
	go build -o ./dist/main
	./dist/main -debug=true -programPath=_resources/Phase1/Prg/4_Colors.prg
