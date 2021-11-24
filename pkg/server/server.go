package server

import (
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/config"
	"github.com/devlibx/gox-base/errors"
	goxServer "github.com/devlibx/gox-base/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sync"
)

type Server interface {
	Start() error
	Stop() chan bool
	GetRouter() *gin.Engine
}

type server struct {
	router         *gin.Engine
	initOnce       *sync.Once
	cf             gox.CrossFunction
	internalServer goxServer.Server
	appConfig      config.App
	logger         *zap.Logger
	startWg        *sync.WaitGroup
}

func (s *server) Start() error {
	var err error

	s.initOnce.Do(func() {

		// Create a new server
		if s.internalServer, err = goxServer.NewServer(s.cf); err != nil {
			err = errors.Wrap(err, "failed to build server")
			return
		}

		// Server is ready - we can stop from this point now
		s.startWg.Done()

		// Start the server
		if err = s.internalServer.Start(s.router, &s.appConfig); err != nil {
			err = errors.Wrap(err, "failed to run http server")
			return
		}
	})

	return err
}

func (s *server) Stop() chan bool {
	s.startWg.Wait()
	return s.internalServer.Stop()
}

func (s *server) GetRouter() *gin.Engine {
	return s.router
}

func NewServer(cf gox.CrossFunction, appConfig config.App) (Server, error) {
	s := server{
		router:    gin.New(),
		cf:        cf,
		initOnce:  &sync.Once{},
		appConfig: appConfig,
		logger:    cf.Logger().Named("server"),
		startWg:   &sync.WaitGroup{},
	}
	s.startWg.Add(1)
	return &s, nil
}
