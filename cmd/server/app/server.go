package app

import (
	"fmt"
	common "github.com/harishb2k/go-template-project/pkg/core/service"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewServerCommand() *cobra.Command {

	app := fx.New(
		fx.Provide(common.NewHelper),
		fx.Provide(NewEntryPoint),
		fx.Invoke(Register),
	)

	return &cobra.Command{
		Use:   "gox",
		Short: "Small help code",
		Long:  `Long help code`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Your service running code here...")
			app.Run()
		},
	}
}
