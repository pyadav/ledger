.PHONY: test

test:
	@go test -race $(shell go list ./...) -coverprofile=coverage.out -count 3

coverage:
	@go test -coverprofile=coverage.out ./... ; go tool cover -func=coverage.out