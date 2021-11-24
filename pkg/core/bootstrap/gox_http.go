package bootstrap

import (
	"github.com/devlibx/gox-base"
	goxHttpApi "github.com/devlibx/gox-http/api"
	"github.com/devlibx/gox-http/command"
)

func NewGoxHttpBuilder(cf gox.CrossFunction, config *command.Config) (goxHttpApi.GoxHttpContext, error) {
	return goxHttpApi.NewGoxHttpContext(cf, config)
}
