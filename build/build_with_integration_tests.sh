go mod tidy
go generate ./...
go fmt ./...

OUTPUT=$(aws dynamodb list-tables  | grep users | wc -l | cut -c8-)
if [ ${OUTPUT} == "1" ]; then
  echo "DynamoDB connection is Ok - we can run integration tests"
  export INCLUDE_DYNAMO_TESTS=true
  export INTEGRATION_TESTS=true
else
    echo "****************** No DynamoDB connection - skip integration tests ******************"
fi

go  test ./... -v
