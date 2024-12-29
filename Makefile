build:
	@go build -o bin/todo cmd/main.go

test:
	@go test -v ./...

run:build
	@./bin/todo 