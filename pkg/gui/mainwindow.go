package gui

import (
	"codenudge/pkg/config"
	"codenudge/pkg/utils"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MainWindow struct {
	CodeEditor    *widget.RichText
	ProgressBar   *widget.ProgressBarInfinite
	Console       *widget.RichText
	ConsoleScroll *container.Scroll
	Window        fyne.Window
}

func NewMainWindow(window fyne.Window, cfgMainWindow *config.MainWindow, cfgShortcuts *config.Shortcut) *MainWindow {
	main := &MainWindow{
		Window: window,
	}
	main.initializeComponents(cfgMainWindow, cfgShortcuts)
	return main
}

func (m *MainWindow) initializeComponents(cfgMainWindow *config.MainWindow, cfgShortcuts *config.Shortcut) {
	m.CodeEditor = widget.NewRichText()
	m.CodeEditor.Wrapping = fyne.TextWrapWord
	m.CodeEditor.AppendMarkdown(utils.Greeting)

	codeScroll := container.NewScroll(m.CodeEditor)
	codeScroll.SetMinSize(fyne.NewSize(300, 300))
	codeArea := container.NewStack(codeScroll)

	m.Console = widget.NewRichText()
	m.Console.Wrapping = fyne.TextWrapWord

	m.ConsoleScroll = container.NewScroll(m.Console)
	m.ConsoleScroll.SetMinSize(fyne.NewSize(300, 100))

	bg := canvas.NewRectangle(color.Gray{Y: 128})
	ConsoleWithBG := container.New(layout.NewStackLayout(), bg, m.ConsoleScroll)

	m.ProgressBar = widget.NewProgressBarInfinite()
	m.ProgressBar.Hide()

	info := container.NewHBox()
	for _, keyCfg := range cfgShortcuts.HotkeyConfigs {
		info.Add(widget.NewLabel(keyCfg.Name))
		info.Add(widget.NewSeparator())
	}
	info.Add(layout.NewSpacer())
	info.Add(container.NewVBox(m.ProgressBar))

	bottomBar := container.NewVBox(ConsoleWithBG, info)

	content := container.NewBorder(nil, bottomBar, nil, nil, codeArea)
	m.Window.SetContent(content)
	m.Window.Resize(fyne.NewSize(cfgMainWindow.WindowWidth, cfgMainWindow.WindowHeight))

}

func (m *MainWindow) Show() {
	m.Window.Show()
}
