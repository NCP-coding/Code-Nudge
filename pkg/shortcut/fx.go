package shortcut

import (
	"codenudge/pkg/config"
	"codenudge/pkg/handler"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func FxModule() fx.Option {
	return fx.Module("shortcut",
		fx.Provide(
			New,
		),
	)
}

type Params struct {
	fx.In
	Config  *config.Shortcut
	Handler *handler.Handler
	Logger  *zap.SugaredLogger
}

func New(p Params) *Shortcut {
	shortcut := NewShortcut(p.Config, p.Handler, p.Logger)
	shortcut.Register()
	return shortcut
}
