package main

import (
	"context"
	app2 "github.com/harishb2k/go-template-project/cmd/server/app"
)

func main() {
	ctx := context.Background()
	app2.Main(ctx, "./config/app.yaml")
	<-ctx.Done()
}
