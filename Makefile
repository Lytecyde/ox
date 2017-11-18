
default: fmt
	go build

fmt:
	go fmt

run:
	go run *.go

lint:
	gometalinter .

test:
	go test

coverage:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
