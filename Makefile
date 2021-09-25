BUILD_PREFIX ?= /tmp

all: server client

server:
	CGO_ENABLED=0 go build -o ${BUILD_PREFIX}/server ./server

client:
	CGO_ENABLED=0 go build -o ${BUILD_PREFIX}/client ./client

test:
	go test -count=1 -timeout=10s ./...

docker-build:
	docker build -f client.Dockerfile --tag=souz9-faraway-words-of-wisdom-client .
	docker build -f server.Dockerfile --tag=souz9-faraway-words-of-wisdom-server .

.PHONY: all server client test docker-build
