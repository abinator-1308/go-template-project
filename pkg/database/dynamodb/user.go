package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/session"
	dynamoOrm "github.com/guregu/dynamo"
	"github.com/harishb2k/go-template-project/pkg/core/service"
	"github.com/harishb2k/go-template-project/pkg/database"
)

type userDaoDynamoImpl struct {
	session  *session.Session
	dynamoDb *dynamoOrm.DB
	table    dynamoOrm.Table
}

func (u *userDaoDynamoImpl) Persist(ctx context.Context, user *database.User) error {
	return u.table.Put(dynamoOrm.AWSEncoding(user)).RunWithContext(ctx)
}

func NewUserDao(dynamo *service.Dynamo) (*userDaoDynamoImpl, error) {
	ud := &userDaoDynamoImpl{
		session:  dynamo.Session,
		dynamoDb: dynamo.DynamoDb,
	}
	ud.table = ud.dynamoDb.Table("users")
	return ud, nil
}
