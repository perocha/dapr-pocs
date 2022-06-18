hello:
	echo "Hello, world!"

build:
	go build -o bin/httpsrv ./cmd/main

docker-build:
	docker build --tag httpsrv  .

clean:
	rm -rf bin