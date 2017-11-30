
default: fmt
	go build

fmt:
	go fmt

run:
	go build && ./ox

lint:
	gometalinter .

test:
	go test ./...

coverage:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
