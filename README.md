
#### Running integration tests
Make table in DynamoDB with name "users"
```shell
Partition	id (String)
Sort  	    key (String)
```

and run ``` INCLUDE_DYNAMO_TESTS=true go test ./...```
