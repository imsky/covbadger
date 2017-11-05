all: quality test build

quality:
	gofmt -w *.go
	go tool vet *.go

test:
	go test -coverprofile=coverage
	go run main.go $$(go tool cover -func=coverage | grep total | cut -d$$'\t' -f5 | cut -d'.' -f1) > coverage.svg

build:
	GOOS=darwin GOARCH=386 go build -o covbadger-darwin
	GOOS=linux GOARCH=386 go build -o covbadger-linux
