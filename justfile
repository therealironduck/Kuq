run *cmd:
    go run ./main.go {{ cmd }}

build:
    go build -o bin/kuq ./main.go

test:
    go test -v ./...
    go vet ./...

lint:
    golangci-lint run ./...
