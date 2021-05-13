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

test:
	go test -v -race ./...

tidy:
	go mod tidy

vendor:
	go mod vendor
