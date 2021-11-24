package main

import (
	"github.com/harishb2k/go-template-project/cmd/server/app"
)

func main() {
	command := app.NewServerCommand()
	if err := command.Execute(); err != nil {
		panic(err)
	}
}
