BINARY_NAME=hello.go
FOLDER=./cmd/hello/

hello:
	echo "Hello, world!"

build:
	go build -o bin/main $(FOLDER)/$(BINARY_NAME)

docker-build:
	docker build --tag $(BINARY_NAME) .