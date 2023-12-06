default: build

setup:
	go mod init

format:
	go fmt ./...

build: 
	# go mod init
	go mod tidy
	go build -o server ./server.go
	go build -o client ./client.go

runs: 
	./server

runc:
	./client

.PHONY: clean
clean:
	-rm -f ./server
	-rm -f ./client
