package handler

import (
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/config"
	"github.com/devlibx/gox-base/metrics"
	"github.com/harishb2k/go-template-project/internal/common"
	"github.com/harishb2k/go-template-project/pkg/bootstrap"
	"github.com/harishb2k/go-template-project/pkg/database"
	"github.com/harishb2k/go-template-project/pkg/server"
	"net/http"
	"sync"
	"time"
)

type UserHandler struct {
	appConfig             config.App
	cf                    gox.CrossFunction
	addUserSuccessCounter metrics.Counter
	userDao               common.UserStore
	messagingFactory      bootstrap.MessagingFactory
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
