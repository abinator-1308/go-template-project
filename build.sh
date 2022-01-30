go mod tidy
go generate ./...
go fmt ./...

export INCLUDE_DYNAMO_TESTS=true
export INTEGRATION_TESTS=true
go  test ./... -v
