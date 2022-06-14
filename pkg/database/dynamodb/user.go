package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/session"
	dynamoOrm "github.com/guregu/dynamo"
	"github.com/harishb2k/go-template-project/pkg/common/objects"
	"github.com/pkg/errors"
)

type UserRepository struct {
	session  *session.Session
	dynamoDb *dynamoOrm.DB
	table    dynamoOrm.Table
}

func (u *UserRepository) Persist(ctx context.Context, user *objects.User) error {
	return u.table.Put(dynamoOrm.AWSEncoding(user)).RunWithContext(ctx)
}

func (u *UserRepository) Get(ctx context.Context, user *objects.User) (*objects.User, error) {
	err := u.table.Get("id", user.ID).Range("key", dynamoOrm.Equal, user.Key).OneWithContext(ctx, dynamoOrm.AWSEncoding(user))
	return user, errors.Wrap(err, "error in fetching user")
}

func (u *UserRepository) UpdateName(ctx context.Context, user *objects.User) error {
	return u.table.Update("id", user.ID).
		Range("key", user.Key).
		Set("name", user.Name).
		RunWithContext(ctx)
}

func NewUserRepository(dynamo *Dynamo) (*UserRepository, error) {
	ud := &UserRepository{
		session:  dynamo.Session,
		dynamoDb: dynamo.DynamoDb,
	}
	ud.table = ud.dynamoDb.Table("users")
	return ud, nil
}
