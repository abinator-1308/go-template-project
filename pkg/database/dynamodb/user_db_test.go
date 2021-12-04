package dynamodb

import (
	"context"
	"github.com/devlibx/gox-base"
	"github.com/harishb2k/go-template-project/pkg/core/service"
	"github.com/harishb2k/go-template-project/pkg/database"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"testing"
	"time"
)

func TestUserPersist(t *testing.T) {
	var userDao database.UserDao
	app := fx.New(
		fx.Provide(gox.NewNoOpCrossFunction),
		service.DynamoServiceModule,
		fx.Provide(NewUserDao),
		fx.Populate(&userDao),
		fx.Provide(func(impl *userDaoDynamoImpl) database.UserDao { return impl }),
		fx.Supply(&service.DynamoConfig{
			Region:  "ap-south-1",
			Timeout: 1,
		}),
	)

	ctx, cf := context.WithTimeout(context.TODO(), 5 * time.Second )
	defer cf()
	_ = app.Start(ctx)

	err := userDao.Persist(ctx, &database.User{
		ID:        "1",
		Key:       "harish",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.NoError(t, err)

}
