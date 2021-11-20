package app

import (
	"fmt"
	common "github.com/harishb2k/go-template-project/pkg/core/service"
)

func Register(helper *common.Helper) {
	fmt.Println("Got helper in the handler.Register() ", helper.Name)
}
