package gui

import (
	"codenudge/pkg/config"

	"fyne.io/fyne/v2/app"
	"go.uber.org/fx"
)

func FxModule() fx.Option {
	return fx.Module("gui",
		fx.Provide(
			New,
		),
	)
}

type Params struct {
	fx.In
	FyneApp    *config.FyneApp
	MainWindow *config.MainWindow
	Shortcut   *config.Shortcut
}

func New(p Params) *Gui {
	app := app.NewWithID(p.FyneApp.Id)
	w := app.NewWindow(p.MainWindow.Title)
	mainWindow := NewMainWindow(w, p.MainWindow, p.Shortcut)
	return NewGui(app, mainWindow)
}
