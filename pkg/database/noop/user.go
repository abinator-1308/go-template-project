package noop

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/session"
	dynamoOrm "github.com/guregu/dynamo"
	"github.com/harishb2k/go-template-project/pkg/database"
	"go.uber.org/fx"
)

var NoopServiceModule = fx.Options(
	fx.Provide(NewUserDao),
	fx.Provide(fx.Annotated{Name: "noopImpl", Target: func(impl *userDaoNoopImpl) database.UserDao {
		return impl
	}}),
)

type userDaoNoopImpl struct {
	session  *session.Session
	dynamoDb *dynamoOrm.DB
	table    dynamoOrm.Table
}

func (u *userDaoNoopImpl) Persist(ctx context.Context, user *database.User) error {
	return nil
}

func (u *userDaoNoopImpl) Get(ctx context.Context, user *database.User) error {
	return nil
}

func (u *userDaoNoopImpl) UpdateName(ctx context.Context, user *database.User) error {
	return nil
}

func NewUserDao() (*userDaoNoopImpl, error) {
	ud := &userDaoNoopImpl{}
	return ud, nil
}
