format:
	go fmt ./...

compile:
	go build ./...

dependencies:
	go mod tidy

lint:
	go vet ./...
	staticcheck ./...

test:
	go test -count=1 -cover ./...

commit: dependencies format compile test lint
