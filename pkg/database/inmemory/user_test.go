package immemory

import (
	"context"
	"github.com/harishb2k/go-template-project/pkg/common/objects"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserDynamoOperations(t *testing.T) {
	ur, err := NewUserRepository()
	assert.NoError(t, err)

	err = ur.Persist(context.Background(), &objects.User{
		ID:        "user_id",
		Property:  "user_key",
		Name:      "user_name",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.NoError(t, err)

	fromDb, err := ur.Get(context.Background(), &objects.User{
		ID:       "user_id",
		Property: "user_key",
	})
	assert.NoError(t, err)
	assert.Equal(t, "user_name", fromDb.Name)

}
