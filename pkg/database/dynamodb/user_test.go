package dynamodb

import (
	"context"
	"github.com/google/uuid"
	"github.com/harishb2k/go-template-project/pkg/database"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"os"
	"testing"
	"time"
)

func TestUserDynamoOperations(t *testing.T) {
	if os.Getenv("INCLUDE_DYNAMO_TESTS") != "true" {
		t.Skip("Skipping integration tests")
	}

	dc := &DynamoConfig{Region: "ap-south-1"}
	var dynamo *Dynamo
	app := fx.New(
		DatabaseModule,
		fx.Supply(dc),
		fx.Populate(&dynamo),
	)
	err := app.Start(context.Background())
	assert.NoError(t, err, "failed to setup app")

	userId := uuid.NewString()
	userDao, err := NewUserRepository(dynamo)
	assert.NoError(t, err, "failed to setup table")

	err = userDao.Persist(context.Background(), &database.User{
		ID:        userId,
		Key:       "harish",
		Name:      "name_1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.NoError(t, err, "failed to save to db")

	fromDb, err := userDao.Get(context.Background(), &database.User{ID: userId, Key: "harish"})
	assert.NoError(t, err, "failed to get from db")
	assert.Equal(t, "name_1", fromDb.Name)
}
