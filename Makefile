vendor:
	go mod tidy -v
	go mod download
	go mod vendor

.PHONY: test
test: FLAGS ?= -parallel 3
test:
	go test -race -covermode=atomic -v ./...

coverage.out:
	go test -v -covermode=count -coverprofile=coverage.out ./...

.PHONY: lint
lint:
	golangci-lint run

go.list:
	go list -json -m all > go.list