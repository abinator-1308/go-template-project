package app

import (
	"context"
	"fmt"
	common "github.com/harishb2k/go-template-project/pkg/core/service"
	"go.uber.org/fx"
)

func NewEntryPoint(lc fx.Lifecycle, helper *common.Helper) string {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting HTTP server. ", helper.Name)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping HTTP server.", helper.Name)
			return nil
		},
	})
	return ""
}
