package command

import (
	"context"
	"fmt"
	"github.com/devlibx/gox-base/serialization"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	app2 "github.com/harishb2k/go-template-project/cmd/server/command"
	"github.com/harishb2k/go-template-project/config"
	"github.com/harishb2k/go-template-project/pkg/database"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	if os.Getenv("INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration tests")
	}

	userId := uuid.NewString()
	key := uuid.NewString()

	// Run the main server
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	appConfig := app2.MainWithConfigAsString(ctx, config.ApplicationConfigString)

	// Setup - call server with
	client := resty.New()
	client.SetHostURL(fmt.Sprintf("http://localhost:%d/%s", appConfig.App.HttpPort, appConfig.App.AppName))

	// First store the new row
	payload := fmt.Sprintf(`{"id": "%s", "key": "%s", "name": "user_random_1"}`, userId, key)
	response, err := client.R().SetBody(payload).Post("/v1/users/")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())

	// Get the data
	response, err = client.R().Get("/v1/users/" + userId + "/" + key)
	assert.NoError(t, err)

	// Check data coming from server
	userFormServer := &database.User{}
	err = serialization.JsonBytesToObject(response.Body(), userFormServer)
	assert.NoError(t, err)
	assert.Equal(t, userId, userFormServer.ID)
	assert.Equal(t, "user_random_1", userFormServer.Name)

	// We are good - lets close it
	cancelFunc()
	<-ctx.Done()
}
