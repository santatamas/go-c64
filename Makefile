.PHONY: build run

build:
	go build -o ./dist/main

run-snake:
	go build -o ./dist/main
	./dist/main -prg-path=_resources/Phase2/7_Snake.prg -no-logs=true

run-stack:
	go build -o ./dist/main
	./dist/main -prg-path=_resources/Phase2/6_stackhack.prg

run:
	go build -o ./dist/main
	./dist/main -prg-path=_resources/Phase1/Prg/4_Colors.prg

debug:
	go build -o ./dist/main
	./dist/main -debug=true -prg-path=_resources/Phase1/Prg/4_Colors.prg
