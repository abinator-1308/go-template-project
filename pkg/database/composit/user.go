package noop

import (
	"context"
	"github.com/harishb2k/go-template-project/pkg/database"
	"go.uber.org/fx"
)

// CompositeDatabaseModule provides works with more than 1 DAO layer. This allows to have more than one data source
// and allows separating read and write to different data sources.
// In this example we have DynamoDb and noop data sources configured. You can perform dual write, read from one and
// write to other etc.
var CompositeDatabaseModule = fx.Options(
	fx.Provide(newUserDao),
	fx.Provide(func(impl *userDaoCompositeImpl) database.UserDao { return impl }),
)

// input is a wrapper which is used to fill dependency my Fx module. Fx module will fill all dependencies here.
type input struct {
	fx.In
	Dynamo database.UserDao `name:"dynamoImpl"`
	Noop   database.UserDao `name:"noopImpl"`
}

type userDaoCompositeImpl struct {
	dynamo database.UserDao
	noop   database.UserDao
}

func (u *userDaoCompositeImpl) Persist(ctx context.Context, user *database.User) error {
	var err error
	if err = u.dynamo.Persist(ctx, user); err == nil {
		err = u.noop.Persist(ctx, user)
	}
	return err
}

func (u *userDaoCompositeImpl) Get(ctx context.Context, user *database.User) error {
	var err error
	if err = u.dynamo.Get(ctx, user); err == nil {
		err = u.noop.Get(ctx, user)
	}
	return err
}

func (u *userDaoCompositeImpl) UpdateName(ctx context.Context, user *database.User) error {
	var err error
	if err = u.dynamo.UpdateName(ctx, user); err == nil {
		err = u.noop.UpdateName(ctx, user)
	}
	return err
}

func newUserDao(input input) (*userDaoCompositeImpl, error) {
	ud := &userDaoCompositeImpl{dynamo: input.Dynamo, noop: input.Noop}
	return ud, nil
}
