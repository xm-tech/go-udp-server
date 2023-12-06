default: build

setup:
	go mod init

format:
	go fmt ./...

build: 
	# go mod init
	go mod tidy
	go build -o boot-server ./server/server.go
	go build -o boot-client ./client/client.go

runs: 
	./boot-server

runc:
	./boot-client

.PHONY: clean
clean:
	-rm -f ./boot-server
	-rm -f ./boot-client
