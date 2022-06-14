package mysql

import (
	"context"
	"github.com/google/uuid"
	"github.com/harishb2k/go-template-project/pkg/common/objects"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"os"
	"testing"
	"time"
)

func TestUserDynamoOperations(t *testing.T) {
	if os.Getenv("INCLUDE_MYSQL_TESTS") != "true" {
		t.Skip("Skipping integration tests")
	}

	var ur *UserRepository
	_ = fx.New(
		fx.Supply(&MySQLConfig{User: "root", Password: "root"}),
		DatabaseModule,
		fx.Populate(&ur),
	).Start(context.Background())

	id := uuid.NewString()
	err := ur.Persist(context.Background(), &objects.User{
		ID:        id,
		Key:       "user_key",
		Name:      "user_name",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.NoError(t, err)

	fromDb, err := ur.Get(context.Background(), &objects.User{
		ID:   id,
		Name: "user_name",
	})
	assert.NoError(t, err)
	assert.Equal(t, "user_name", fromDb.Name)

}
