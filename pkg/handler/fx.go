package handler

import (
	"codenudge/pkg/clients"
	"codenudge/pkg/gui"
	"codenudge/pkg/screenshotter"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func FxModule() fx.Option {
	return fx.Module("handler",
		fx.Provide(
			New,
		),
	)
}

type Params struct {
	fx.In
	Gui           *gui.Gui
	GenAiClient   *clients.GenAiClient
	Screenshotter *screenshotter.Screenshotter
	Logger        *zap.SugaredLogger
}

func New(p Params) *Handler {
	handler := NewHandler(p.Gui, p.GenAiClient, p.Screenshotter, p.Logger)
	return handler
}
