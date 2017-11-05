
default: fmt
	go build

fmt:
	go fmt

run:
	go run *.go

lint:
	gometalinter .