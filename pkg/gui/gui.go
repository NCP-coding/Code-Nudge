package gui

import "fyne.io/fyne/v2"

type Gui struct {
	App        fyne.App
	MainWindow *MainWindow
}

func NewGui(app fyne.App, mainwindow *MainWindow) *Gui {
	return &Gui{
		App:        app,
		MainWindow: mainwindow,
	}
}

func (g *Gui) ShowAndRun() {
	g.MainWindow.Show()
	g.App.Run()
}
