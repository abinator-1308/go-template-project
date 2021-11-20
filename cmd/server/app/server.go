package app

import (
	"context"
	"fmt"
	common2 "github.com/harishb2k/go-template-project/pkg/common"
	common "github.com/harishb2k/go-template-project/pkg/core/service"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"time"
)

func NewServerCommand() *cobra.Command {

	app := fx.New(
		fx.Provide(common.NewHelper), // NewHelper will be called only if someone needs "*Helper" object.
		fx.Invoke(Register),          // This will be called by app.Start()
		fx.Invoke(NewEntryPoint),     // This will be called by app.Start()
		fx.Supply(110),
		fx.Supply(&common2.Config{Env: "prod", Port: 11}), // Everyone will have access to this object
	)

	return &cobra.Command{
		Use:   "gox",
		Short: "Small help code",
		Long:  `Long help code`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Your service running code here...")
			err := app.Start(context.TODO())
			if err != nil {
				fmt.Println("Got error", err)
			}
			fmt.Println("Service done")
			time.Sleep(2 * time.Second)
		},
	}
}
