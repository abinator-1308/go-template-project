package app

import (
	"context"
	"github.com/devlibx/gox-base/serialization"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	app2 "github.com/harishb2k/go-template-project/cmd/server/app"
	"github.com/harishb2k/go-template-project/pkg/database"
	"github.com/harishb2k/go-template-project/pkg/database/dynamodb"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"os"
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	if os.Getenv("INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration tests")
	}

	dc := &dynamodb.DynamoConfig{Region: "ap-south-1"}
	var dynamo *dynamodb.Dynamo
	var ur *dynamodb.UserRepository
	app := fx.New(
		dynamodb.DatabaseModule,
		fx.Supply(dc),
		fx.Populate(&dynamo, &ur),
	)
	err := app.Start(context.Background())
	assert.NoError(t, err, "failed to setup app")

	userId := uuid.NewString()
	key := uuid.NewString()
	assert.NoError(t, err, "failed to setup table")

	err = ur.Persist(context.Background(), &database.User{
		ID:        userId,
		Key:       key,
		Name:      "name_1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.NoError(t, err, "failed to save to db")

	ctx := context.Background()
	ctxWithTimeout, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	app2.Main(ctxWithTimeout, "../../../../../config/app.yaml")

	client := resty.New()
	client.SetHostURL("http://localhost:8090/ser")

	response, err := client.R().
		Get("/v1/users/" + userId + "/" + key)
	assert.NoError(t, err)

	userFormServer := &database.User{}
	err = serialization.JsonBytesToObject(response.Body(), userFormServer)
	assert.NoError(t, err)
	assert.Equal(t, userId, userFormServer.ID)
	assert.Equal(t, "name_1", userFormServer.Name)

	cancelFunc()
	<-ctxWithTimeout.Done()
}
