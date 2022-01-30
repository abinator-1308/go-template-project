package main

import (
	"context"
	"github.com/harishb2k/go-template-project/cmd/server/command"
)

func main() {
	ctx := context.Background()
	command.Main(ctx, "./config/app.yaml")
	<-ctx.Done()
}
