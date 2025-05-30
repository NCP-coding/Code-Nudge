package logger

import (
	"codenudge/pkg/gui"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func FxModule() fx.Option {
	return fx.Module("logger",
		fx.Provide(
			NewLogger,
		),
	)
}

type Params struct {
	fx.In
	Gui *gui.Gui
}

func NewLogger(p Params) *zap.SugaredLogger {
	sink := NewFynWidgetSink(p.Gui)
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(sink), zapcore.InfoLevel)
	return zap.New(core).Sugar()
}
