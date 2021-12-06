package noop

import (
	"context"
	"github.com/harishb2k/go-template-project/pkg/database"
	"go.uber.org/fx"
)

var DatabaseModule = fx.Options(
	fx.Provide(newUserDao),
	fx.Provide(fx.Annotated{Name: "noopImpl", Target: func(impl *userDaoNoopImpl) database.UserDao {
		return impl
	}}),
)

type userDaoNoopImpl struct {
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

func newUserDao() (*userDaoNoopImpl, error) {
	ud := &userDaoNoopImpl{}
	return ud, nil
}
