cover:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func coverage.out

coverhtml:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

dep:
	go mod download

format:
	gofmt -s -w .

lint:
	golangci-lint run ./...

test:
	go test -v -race ./...

tidy:
	go mod tidy

vendor:
	go mod vendor
