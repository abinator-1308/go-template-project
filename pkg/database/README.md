#### How to set up repository

A repository can be setup with following example. This code will populate ```userRepository``` with
```composite.UserRepository```.

```golang
package "main"

import (
"github.com/harishb2k/go-template-project/pkg/database/composite"
)

var userRepository *composite.UserRepository
app := fx.New(
    fx.Provide(gox.NewNoOpCrossFunction),
    CompositeDatabaseModule,
    fx.Populate(&userRepository),
    fx.Supply(&dynamodb.DynamoConfig{Region: "ap-south-1", Timeout: 1}),
)
```

Note we have created a "composite" with a dynamo db + noop impl.

#### Enable MySQL support

1. To enable tests create db names "test" with "root/root", from file 'pkg/database/mysql/schema/schema.sql'
2. Run tests ```INCLUDE_MYSQL_TESTS=true go test ./... ```

#### Running integration tests

```shell

# 1. MySQL 
# Make sure you have a db created "test":
# Create tables form schema.sql file
# Make a user root/root with all access

export INCLUDE_DYNAMO_TESTS=false
export INTEGRATION_TESTS=true
export INCLUDE_MYSQL_TESTS=true
go  test ./... -v
```

### Working with MySQL
Install SqlC tool which we use to generate SQL code
```shell
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

Go to `/pkg/database/mysql/schema` and add schema and new queries

Generate code with latest queries
```shell
cd pkg/database/mysql/schema
sqlc generate
```