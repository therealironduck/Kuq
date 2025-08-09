run cmd:
    go run ./cmd/kuq/main.go {{ cmd }}

build:
    go build -o bin/kuq ./cmd/kuq/main.go

test:
    go test -v ./...
    go vet ./...

lint:
    golangci-lint run ./...
