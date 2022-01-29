package noop

import (
	"context"
	"github.com/harishb2k/go-template-project/pkg/database"
	"go.uber.org/fx"
)

var DatabaseModule = fx.Options(
	fx.Provide(NewUserRepository),
)

type UserRepository struct {
}

func (u *UserRepository) Persist(ctx context.Context, user *database.User) error {
	return nil
}

func (u *UserRepository) Get(ctx context.Context, user *database.User) (*database.User, error) {
	return nil, nil
}

func (u *UserRepository) UpdateName(ctx context.Context, user *database.User) error {
	return nil
}

func NewUserRepository() (*UserRepository, error) {
	ud := &UserRepository{}
	return ud, nil
}
