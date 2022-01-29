#### 
```golang
import(
    "github.com/harishb2k/go-template-project/pkg/database/composit"
)

var userDao *composit.UserRepository
app := fx.New(
    fx.Provide(gox.NewNoOpCrossFunction),
    CompositeDatabaseModule,
    fx.Populate(&userDao),
    fx.Supply(&dynamodb.DynamoConfig{Region: "ap-south-1", Timeout: 1}),
)
```