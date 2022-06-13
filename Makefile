hello:
	echo "Hello, world!"

build-hello:
	go build -o bin/main ./cmd/hello

build-http-srv:
	go build -o bin/http-srv ./cmd/http-srv

docker-build-hello:
	docker build --tag hello --build-arg MODULE_NAME=hello  .

docker-build-http-srv:
	docker build --tag http-srv --build-arg MODULE_NAME=http-srv  .

clean:
	rm -rf bin