
test:
	@go test -race ./...

clean:
	@go clean
	@find . -type f -name "mock_*.go" -delete
