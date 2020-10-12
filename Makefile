.PHONY: all build vet test clean

all: clean build vet test

build: clean
	go build -o kv-store main.go

vet:
	go vet ./...

test: build
	go test ./...

clean:
	rm -f kv-store
