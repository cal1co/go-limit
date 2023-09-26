build:
	go build -o bin/go-limiter

run: build
	./bin/go-limiter

test:
	go test -v ./...