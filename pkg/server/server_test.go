package server

import (
	"context"
	"errors"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	// Setup a server
	var s Server
	app := fx.New(
		fx.Provide(gox.NewNoOpCrossFunction),
		fx.Supply(config.App{
			AppName:     "test",
			HttpPort:    18901,
			Environment: "test",
		}),
		fx.Provide(NewServer),
		fx.Populate(&s),
	)
	defer s.Stop()

	ctx, ch := context.WithTimeout(context.TODO(), 5*time.Second)
	defer ch()

	// Start this app
	if err := app.Start(ctx); err != nil {
		assert.NoError(t, err)
	}
	defer app.Stop(ctx)

	// Start server and wait for 1 sec to boot the server
	go func() {
		err := s.Start()
		assert.True(t, errors.Is(err, http.ErrServerClosed))
	}()
	time.Sleep(1 * time.Second)

	select {
	case <-time.After(5 * time.Second):
		t.Fail()
	case <-s.Stop():
	}
}

// Sample is an example of a server
// What is enabled:
// 1. Logging
// 2. Graceful Shutdown
// 3. Gin router
// 4. Panic recovery
func Sample() {
	server, err := NewServer(gox.NewNoOpCrossFunction(), config.App{
		AppName:     "example",
		HttpPort:    18090,
		Environment: "test",
	})
	if err != nil {
		panic(err)
	}

	go func() {
		if err := server.Start(); err != nil {
			// Note - http.ErrServerClosed error is normal which happens when you stop (you can handle it you want)
			panic(err)
		}
	}()
	time.Sleep(1 * time.Second)

	// Stop server when you are done
	<-server.Stop()
}
