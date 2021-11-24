package app

import (
	"context"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/config"
	"github.com/golang/mock/gomock"
	"github.com/harishb2k/go-template-project/internal/handler"
	"github.com/harishb2k/go-template-project/pkg/clients/jsonplaceholder"
	mockJsonplaceholderClient "github.com/harishb2k/go-template-project/pkg/mocks/clients/jsonplaceholder"
	"github.com/harishb2k/go-template-project/pkg/server"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"net/http"
	"net/http/httptest"
	"testing"
)

// End-to-End test to test out setup
func TestServer(t *testing.T) {
	ctrl := gomock.NewController(t)

	var s ServerImpl
	app := fx.New(

		// Register all HTTP API handlers
		handler.RandomHandlerModule,
		handler.UserHandlerModule,

		// Basic dependency - underlying server, CrossFunc, configs for application
		fx.Provide(server.NewServer),
		fx.Provide(gox.NewNoOpCrossFunction),
		fx.Supply(config.App{}),

		// Instance of underlying server
		fx.Populate(&s),

		// Provide mocks
		fx.Provide(func() jsonplaceholder.Client { return mockJsonplaceholderClient.NewMockClient(ctrl) }),
	)
	_ = app.Start(context.TODO())
	// defer app.Stop(context.TODO())
	// defer s.Stop()

	r := httptest.NewRequest("POST", "/srv/v1/users", nil)
	w := httptest.NewRecorder()
	s.routes()
	s.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
