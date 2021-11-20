package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "gox",
		Short: "Small help code",
		Long:  `Long help code`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Your service running code here...")
		},
	}
}
