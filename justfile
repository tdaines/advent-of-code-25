default: run

test:
    go test -v ./...

run:
    go run .

build:
    go build -v main.go
