all: quality test build

quality:
	gofmt -w *.go
	go tool vet *.go

test:
	go test -cover

build:
	GOOS=darwin GOARCH=386 go build -o covbadger-darwin
	GOOS=linux GOARCH=386 go build -o covbadger-linux
