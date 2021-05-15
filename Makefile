build:
	go build -o bin/main cmd/api/main.go

cover:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func coverage.out

coverhtml:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

dep:
	go mod download

lint:
	golangci-lint run ./...

pretty:
	gofmt -s -w .

run:
	go run cmd/api/main.go

test:
	go test -v -race ./...

tidy:
	go mod tidy

vendor:
	go mod vendor
