package main

import (
	"fmt"
	"github.com/harishb2k/go-template-project/cmd/server/app"
	"os"
)

func main() {
	command := app.NewServerCommand()
	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
