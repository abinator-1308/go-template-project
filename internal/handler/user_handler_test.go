package handler

import (
	"context"
	"fmt"
	"github.com/devlibx/gox-base/serialization"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/harishb2k/go-template-project/pkg/common/objects"
	"github.com/harishb2k/go-template-project/pkg/server"
	commonTesting "github.com/harishb2k/go-template-project/pkg/testing"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddUser(t *testing.T) {
	// Get the user handler
	var uh *UserHandler
	var r *gin.Engine
	err := fx.New(
		commonTesting.TestCommonModule,
		TestUserHandlerModule,
		fx.Populate(&uh, &r),
	).Start(context.Background())
	assert.NoError(t, err)

	// Setup dummy end-point to test handler
	id := uuid.NewString()
	key := uuid.NewString()

	r.POST("/users", func(c *gin.Context) {
		server.EnsureGinContextWrapper(uh.Adduser()).ServeHTTP(c.Writer, c.Request)
	})
	r.GET("/users/:id/:property", func(c *gin.Context) {
		server.EnsureGinContextWrapper(uh.GetUser()).ServeHTTP(c.Writer, c.Request)
	})

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"id": "`+id+`", "property": "`+key+`", "name": "3"}`))
	w := httptest.NewRecorder()

	// Trigger serve api to test end to end
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Make a get request and see everything is working
	req = httptest.NewRequest(http.MethodGet, "/users/"+id+"/"+key, nil)
	w = httptest.NewRecorder()

	// Trigger serve api to test end to end
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body)
	u := &objects.User{}
	err = serialization.JsonBytesToObject(w.Body.Bytes(), u)
	assert.NoError(t, err)
	assert.Equal(t, id, u.ID)
	assert.Equal(t, key, u.Property)
	assert.Equal(t, "3", u.Name)
}
