package handler

import (
	"github.com/devlibx/gox-base"
	"github.com/gin-gonic/gin"
	immemory "github.com/harishb2k/go-template-project/pkg/database/inmemory"
	"github.com/harishb2k/go-template-project/pkg/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddUser(t *testing.T) {
	userDao, err := immemory.NewUserRepository()
	assert.NoError(t, err)

	// Setup a user handler
	uh := UserHandler{
		cf:      gox.NewNoOpCrossFunction(),
		userDao: userDao,
	}

	// Setup a Gin router to test our API
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(server.GinContextToContextMiddleware())
	r.POST("/users", func(c *gin.Context) {
		server.EnsureGinContextWrapper(uh.Adduser()).ServeHTTP(c.Writer, c.Request)
	})

	// Make a dummy request and see everything is working
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"id": "1", "key": "2", "name": "3"}`))
	w := httptest.NewRecorder()

	// Trigger serve api to test end to end
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
