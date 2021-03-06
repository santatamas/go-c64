.PHONY: build run

build:
	go build -o ./dist/main

run-colors:
	go build -o ./dist/main
	./dist/main -prg-path=_resources/prg/4_colors.prg

run-addr:
	go build -o ./dist/main
	./dist/main -prg-path=_resources/prg/5_addressingmodes.prg

run-monopole:
	go build -o ./dist/main
	./dist/main -prg-path=_resources/prg/monopole.prg

run-snake:
	go build -o ./dist/main
	./dist/main -prg-path=_resources/prg/7_Snake.prg -no-logs=true

run-stack:
	go build -o ./dist/main
	./dist/main -prg-path=_resources/prg/6_stackhack.prg

run:
	go build -o ./dist/main
	./dist/main -no-logs=true

run-log:
	go build -o ./dist/main
	./dist/main

run-bin:
	go build -o ./dist/main
	./dist/main -prg-path=_resources/rom/64c.251913-01.bin

debug-client:
	cd internals/WebUI; ng serve --open

debug:
	go build -o ./dist/main
	./dist/main -debug=true

debug-test-suite:
	go build -o ./dist/main
	./dist/main -debug=true -test=true

tail:
	tail -f log.txt

debug-first:
	go build -o ./dist/main
	./dist/main -debug=true -prg-path=_resources/prg/1_First.prg
