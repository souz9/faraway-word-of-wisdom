BUILD_PREFIX ?= /tmp

all: server client

server:
	go build -o ${BUILD_PREFIX}/server ./server

client:
	go build -o ${BUILD_PREFIX}/client ./client

test:
	go test -count=1 -timeout=10s ./...

.PHONY: all server client test
