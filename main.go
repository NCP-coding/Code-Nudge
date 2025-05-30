package main

import (
	"context"

	"go.uber.org/fx"

	"codenudge/pkg/clients"
	"codenudge/pkg/config"
	"codenudge/pkg/gui"
	"codenudge/pkg/handler"
	"codenudge/pkg/logger"
	"codenudge/pkg/screenshotter"
	"codenudge/pkg/shortcut"
)

func main() {
	var ui *gui.Gui

	app := fx.New(
		fx.Provide(context.Background),
		config.FxModule(),
		logger.FxModule(),
		clients.FxModule(),
		gui.FxModule(),
		screenshotter.FxModule(),
		handler.FxModule(),
		shortcut.FxModule(),
		fx.Populate(&ui),
		fx.Invoke(shortcut.Run),
	)
	go app.Run()

	ui.ShowAndRun()
}
