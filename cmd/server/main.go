package main

import (
	"context"
	"github.com/devlibx/gox-base/serialization"
	"github.com/harishb2k/go-template-project/cmd/server/command"
	"github.com/harishb2k/go-template-project/internal/config"
)

func main() {

	// Read config
	appConfig := config.ApplicationConfig{}
	err := serialization.ReadYaml("./config/app.yaml", &appConfig)
	if err != nil {
		panic(err)
	}

	// Run the main command to launch application
	ctx := context.Background()
	command.Main(ctx, &appConfig)
	<-ctx.Done()
}
