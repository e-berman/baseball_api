build:
	go build -o bin/baseball_api

run: build
	./bin/baseball_api

test:
	go test -v ./...
