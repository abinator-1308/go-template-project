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

var IntegrationJsonPlaceholderModule = fx.Options(
	fx.Provide(jsonplaceholder.NewJsonPlaceHolderClient),
	fx.Provide(func(client *jsonplaceholder.Client) JsonPlaceholderClient { return client }),
)
