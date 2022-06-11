BINARY_NAME=./cmd/hello/hello.go

hello:
	echo "Hello, world!"

build:
	go build -o bin/main $(BINARY_NAME)