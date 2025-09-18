.PHONY: build run

build:
	go build -o notas .

run: build
	./notas
