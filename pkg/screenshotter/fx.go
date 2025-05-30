package screenshotter

import (
	"codenudge/pkg/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func FxModule() fx.Option {
	return fx.Module("screenshotter",
		fx.Provide(
			New,
		),
	)
}

type Params struct {
	fx.In
	Config *config.Screenshot
	Logger *zap.SugaredLogger
}

func New(p Params) *Screenshotter {
	screenshotter := NewScreenshotter(p.Config, p.Logger)
	return screenshotter
}
