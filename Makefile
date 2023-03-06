test:
	go test -v -short ./...

e2e:
	go test -v ./internal/parser...

build:
	go build ./...

test-e2e:
	dlv test github.com/m4tthewde/gov1/internal/parser
