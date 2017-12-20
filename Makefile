help:
	cat Makefile

build:
	go build server.go

run:
	make build
	./server
