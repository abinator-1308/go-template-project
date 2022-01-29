package composit

import (
	"context"
	"github.com/harishb2k/go-template-project/pkg/database"
	"github.com/harishb2k/go-template-project/pkg/database/dynamodb"
	"github.com/harishb2k/go-template-project/pkg/database/noop"
	"go.uber.org/fx"
)

// CompositeDatabaseModule provides works with more than 1 DAO layer. This allows to have more than one data source
// and allows separating read and write to different data sources.
// In this example we have DynamoDb and noop data sources configured. You can perform dual write, read from one and
// write to other etc.
var CompositeDatabaseModule = fx.Options(
	fx.Provide(NewUserRepository),
	// fx.Provide(func(impl *UserRepository) *UserRepository { return impl }),
	dynamodb.DatabaseModule,
	noop.DatabaseModule,
)

// input is a wrapper which is used to fill dependency my Fx module. Fx module will fill all dependencies here.
type input struct {
	fx.In
	Dynamo *dynamodb.UserRepository `name:"dynamoImpl"`
	Noop   *noop.UserRepository     `name:"noopImpl"`
}

type UserRepository struct {
	dynamo *dynamodb.UserRepository
	noop   *noop.UserRepository
}

func (u *UserRepository) Persist(ctx context.Context, user *database.User) error {
	var err error
	if err = u.dynamo.Persist(ctx, user); err == nil {
		err = u.noop.Persist(ctx, user)
	}
	return err
}

func (u *UserRepository) Get(ctx context.Context, user *database.User) (*database.User, error) {
	var err error
	var found *database.User
	if found, err = u.dynamo.Get(ctx, user); err == nil {
		return found, nil
	}
	if found, err = u.noop.Get(ctx, user); err == nil {
		return found, err
	}
	return nil, err
}

func (u *UserRepository) UpdateName(ctx context.Context, user *database.User) error {
	var err error
	if err = u.dynamo.UpdateName(ctx, user); err == nil {
		err = u.noop.UpdateName(ctx, user)
	}
	return err
}

func NewUserRepository(input input) (*UserRepository, error) {
	ud := &UserRepository{dynamo: input.Dynamo, noop: input.Noop}
	return ud, nil
}
