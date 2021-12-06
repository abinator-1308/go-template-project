package noop

import (
	"context"
	"github.com/harishb2k/go-template-project/pkg/database"
	"go.uber.org/fx"
)

// CompositeHandlerModule provides example of a random API you can add
var CompositeHandlerModule = fx.Options(
	fx.Provide(NewUserDao),
	fx.Provide(func(impl *userDaoCompositeImpl) database.UserDao { return impl }),
)

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
	if err := u.dynamo.Persist(ctx, user); err == nil {
		if err := u.noop.Persist(ctx, user); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func (u *userDaoCompositeImpl) Get(ctx context.Context, user *database.User) error {
	if err := u.dynamo.Get(ctx, user); err == nil {
		if err := u.noop.Get(ctx, user); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func (u *userDaoCompositeImpl) UpdateName(ctx context.Context, user *database.User) error {
	if err := u.dynamo.UpdateName(ctx, user); err == nil {
		if err := u.noop.UpdateName(ctx, user); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func NewUserDao(input input) (*userDaoCompositeImpl, error) {
	ud := &userDaoCompositeImpl{dynamo: input.Dynamo, noop: input.Noop}
	return ud, nil
}
