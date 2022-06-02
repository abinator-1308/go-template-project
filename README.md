
#### Running integration tests
Make table in DynamoDB with name "users"
```shell
Partition	id (String)
Sort  	    key (String)
```

and run 
```shell
# You should have DynamoDB access
# you should have kafka running  

export INCLUDE_DYNAMO_TESTS=true
export INTEGRATION_TESTS=true
go  test ./... -v
```

---
## Add MySQL support
Go to ```pkg/database/README.md```