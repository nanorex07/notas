.PHONY: build run

build:
	go clean && go build -o build/notas .

run:
	./build/notas
