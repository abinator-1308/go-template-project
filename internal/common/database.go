package common

import (
	"context"
	"github.com/harishb2k/go-template-project/pkg/common/objects"
)

type UserStore interface {
	Persist(ctx context.Context, user *objects.User) error
	Get(ctx context.Context, user *objects.User) (*objects.User, error)
}
