default: build

setup:
	go mod init

format:
	go fmt ./...

build: 
	# go mod init
	go mod tidy
	go build -o server ./server/server.go
	go build -o client ./client/client.go

runs: 
	./server/server

runc:
	./client/client

.PHONY: clean
clean:
	-rm -f ./server/server
	-rm -f ./client/client
