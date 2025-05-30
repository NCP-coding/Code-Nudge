package clients

import (
	"context"
	"codenudge/pkg/config"

	"go.uber.org/fx"
)

func FxModule() fx.Option {
	return fx.Module("genaiclient",
		fx.Provide(
			New,
		),
	)
}

type Params struct {
	fx.In
	Config  *config.GenaiClient
	Context context.Context
}

func New(p Params) *GenAiClient {
	return NewGenAiClient(p.Config, p.Context)
}
