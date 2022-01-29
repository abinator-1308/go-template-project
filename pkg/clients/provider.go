package clients

import (
	"github.com/harishb2k/go-template-project/pkg/clients/jsonplaceholder"
	"go.uber.org/fx"
)

// Module defines all the clients available to caller
// 1. Example Json Placeholder
// 2. ...
var Module = fx.Options(
	fx.Provide(jsonplaceholder.NewJsonPlaceHolderClient),
)

// IntegrationJsonPlaceholderModule is a completely wired module for ease of use in application. If you want everything
// wired-up automatically then use this
var IntegrationJsonPlaceholderModule = fx.Options(
	fx.Provide(jsonplaceholder.NewJsonPlaceHolderClient),
	fx.Provide(func(client *jsonplaceholder.Client) JsonPlaceholderClient { return client }),
)

// IntegrationModule is a completely wired module for ease of use in application. If you want everything
// wired-up automatically then use this
var IntegrationModule = fx.Options(
	IntegrationJsonPlaceholderModule,
)
