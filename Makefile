all: quality test build

quality:
	gofmt -w *.go
	go tool vet *.go

test:
	go test -cover

build:
	go build
