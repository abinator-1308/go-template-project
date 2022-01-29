package common

import (
	"context"
	"github.com/harishb2k/go-template-project/pkg/database"
)

type UserStore interface {
	Persist(ctx context.Context, user *database.User) error
}
