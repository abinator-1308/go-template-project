#### How to set up repository

A repository can be setup with following example. This code will populate ```userRepository``` with
```composite.UserRepository```.

```golang
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

