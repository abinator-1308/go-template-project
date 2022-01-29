package composit

import (
	"context"
	"fmt"
	"github.com/devlibx/gox-base"
	"github.com/google/uuid"
	"github.com/harishb2k/go-template-project/pkg/database"
	"github.com/harishb2k/go-template-project/pkg/database/dynamodb"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"os"
	"testing"
	"time"
)

func TestUserPersist(t *testing.T) {
	if os.Getenv("INCLUDE_DYNAMO_TESTS") != "true" {
		t.Skip("Skipping integration tests")
	}

	var userDao *UserRepository
	app := fx.New(
		fx.Provide(gox.NewNoOpCrossFunction),
		CompositeDatabaseModule,
		fx.Populate(&userDao),
		fx.Supply(&dynamodb.DynamoConfig{Region: "ap-south-1", Timeout: 1}),
	)

	ctx, cf := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cf()
	_ = app.Start(ctx)

	userId := uuid.NewString()
	fmt.Println("User Id for test = ", userId)
	err := userDao.Persist(ctx, &database.User{
		ID:        userId,
		Key:       "harish",
		Name:      "name_1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.NoError(t, err)

	fromDb := &database.User{ID: userId, Key: "harish"}
	err = userDao.Get(ctx, fromDb)
	assert.NoError(t, err)
	assert.Equal(t, "name_1", fromDb.Name)

	userWithNewName := &database.User{ID: userId, Key: "harish", Name: "name_2"}
	err = userDao.UpdateName(ctx, userWithNewName)
	assert.NoError(t, err)
	err = userDao.Get(ctx, userWithNewName)
	assert.NoError(t, err)
	assert.Equal(t, "name_2", userWithNewName.Name)
}
