package immemory

import (
	"context"
	"github.com/devlibx/gox-base/errors"
	"github.com/harishb2k/go-template-project/pkg/database"
	"go.uber.org/fx"
)

var DatabaseModule = fx.Options(
	fx.Provide(NewUserRepository),
)

type UserRepository struct {
	users map[string]map[string]*database.User
}

func (u *UserRepository) Persist(ctx context.Context, user *database.User) error {
	if temp, ok := u.users[user.ID]; ok {
		temp[user.Key] = user
	} else {
		u.users[user.ID] = map[string]*database.User{}
		u.users[user.ID][user.Key] = user
	}
	return nil
}

func (u *UserRepository) Get(ctx context.Context, user *database.User) (*database.User, error) {
	if temp, ok := u.users[user.ID]; ok {
		if temp1, ok := temp[user.Key]; ok {
			return temp1, nil
		}
	}
	return nil, errors.New("not found")
}

func (u *UserRepository) UpdateName(ctx context.Context, user *database.User) error {
	if temp, ok := u.users[user.ID]; ok {
		if temp1, ok := temp[user.Key]; ok {
			temp1.Name = user.Name
			return nil
		}
	}
	return errors.New("not found")
}

func NewUserRepository() (*UserRepository, error) {
	ud := &UserRepository{
		users: map[string]map[string]*database.User{},
	}
	return ud, nil
}
