package logger

import (
	"codenudge/pkg/gui"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type FyneWidgetSink struct {
	Gui *gui.Gui
	mu  sync.Mutex
}

func NewFynWidgetSink(gui *gui.Gui) *FyneWidgetSink {
	return &FyneWidgetSink{
		Gui: gui,
	}
}

func (fs *FyneWidgetSink) Write(p []byte) (n int, err error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	fyne.Do(func() {
		segments := &fs.Gui.MainWindow.Console.Segments
		*segments = append(*segments, &widget.TextSegment{
			Text: strings.TrimSuffix(string(p), "\n"),
		})
		fs.Gui.MainWindow.Console.Refresh()
		fs.Gui.MainWindow.ConsoleScroll.ScrollToBottom()
	})

	return len(p), nil
}

func (fs *FyneWidgetSink) Sync() error {
	return nil
}
