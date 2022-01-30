package handler

import (
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/config"
	"github.com/devlibx/gox-base/metrics"
	"github.com/harishb2k/go-template-project/internal/common"
	"github.com/harishb2k/go-template-project/pkg/database"
	"github.com/harishb2k/go-template-project/pkg/server"
	"go.uber.org/fx"
	"net/http"
	"sync"
	"time"
)

// UserHandlerModule has all the HTTP hap handlers for user modules. By taking this approach we are able to:
// 1. encapsulate all handlers for user
// 2. User Handler can get everything injected and all handlers can use those dependencies
// 3. Since we return plain HTTP handler, it can be ued by any framework (however you can have Gin specific code here)
var UserHandlerModule = fx.Options(
	fx.Provide(func(cf gox.CrossFunction, appConfig config.App, userDao common.UserStore) *UserHandler {
		return &UserHandler{
			appConfig: appConfig,
			cf:        cf,
			userDao:   userDao,
		}
	}),
)

type UserHandler struct {
	appConfig             config.App
	cf                    gox.CrossFunction
	addUserSuccessCounter metrics.Counter
	userDao               common.UserStore
}

func (uh *UserHandler) Adduser() http.HandlerFunc {
	initOnce := sync.Once{}
	return func(w http.ResponseWriter, r *http.Request) {
		initOnce.Do(func() {
			uh.addUserSuccessCounter = uh.cf.Metric().Counter("handler__add_user_success")
		})

		// Get gin context if you want to use
		ginContext := server.GinContextFromHttpRequestVerified(r)
		_ = ginContext

		// Expected request
		type user struct {
			ID   string `json:"id"`
			Key  string `json:"key"`
			Name string `json:"name"`
		}
		u := &user{}
		err := ginContext.BindJSON(u)

		if err == nil {
			err = uh.userDao.Persist(r.Context(), &database.User{
				ID:        u.ID,
				Key:       u.Key,
				Name:      u.Name,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		// Do your logic
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ok"))
		uh.addUserSuccessCounter.Inc(1)
	}
}

func (uh *UserHandler) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ginContext := server.GinContextFromHttpRequestVerified(r)
		id := ginContext.Param("id")
		key := ginContext.Param("key")

		if user, err := uh.userDao.Get(r.Context(), &database.User{ID: id, Key: key}); err == nil {
			ginContext.JSON(http.StatusOK, user)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
