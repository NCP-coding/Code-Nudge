package handler

import (
	"codenudge/pkg/clients"
	"codenudge/pkg/gui"
	"codenudge/pkg/screenshotter"
	"codenudge/pkg/utils"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"go.uber.org/zap"
)

type Handler struct {
	Gui           *gui.Gui
	GenAiClient   *clients.GenAiClient
	Screenshotter *screenshotter.Screenshotter
	logger        *zap.SugaredLogger
	Callbacks     map[string]func()
}

func NewHandler(gui *gui.Gui, genAiClient *clients.GenAiClient, screenshotter *screenshotter.Screenshotter, logger *zap.SugaredLogger) *Handler {
	handler := &Handler{
		Gui:           gui,
		GenAiClient:   genAiClient,
		Screenshotter: screenshotter,
		logger:        logger,
	}

	handler.Callbacks = map[string]func(){
		"Alt + S: Test UI":           handler.TestProgressBar,
		"Ctrl + Alt + A: Screenshot": handler.CaptureDisplayAndGenAnswer,
		"Alt + T: Example HINT":      handler.ExampleAnswer,
	}
	return handler
}

func (h *Handler) StartProgress() {
	fyne.Do(func() {
		h.Gui.MainWindow.ProgressBar.Show()
		h.Gui.MainWindow.ProgressBar.Start()
	})
}

func (h *Handler) StopProgress() {
	fyne.Do(func() {
		h.Gui.MainWindow.ProgressBar.Stop()
		h.Gui.MainWindow.ProgressBar.Hide()
	})
}

func (h *Handler) UpdateRichText(markdown string) {
	fyne.Do(func() {
		h.Gui.MainWindow.CodeEditor.Segments = []widget.RichTextSegment{}
		h.Gui.MainWindow.CodeEditor.AppendMarkdown(markdown)
	})
}

func (h *Handler) TestProgressBar() {
	h.StartProgress()
	time.Sleep(1 * time.Second)
	h.StopProgress()
}

func (h *Handler) ExampleAnswer() {
	h.UpdateRichText(utils.SubarraySumExplanation)
}

func (h *Handler) CaptureDisplayAndGenAnswer() {
	h.StartProgress()
	defer h.StopProgress()

	buf, err := h.Screenshotter.CaptureAndSafeInDir()
	if err != nil {
		h.logger.Errorf("Failed capturing screenshot due to %v", err)
		return
	}
	h.logger.Info("Uploading file")
	file, restore, err := h.GenAiClient.UploadFile(buf)
	if err != nil {
		h.logger.Errorf("Failed uploading file due to %v", err)
		return
	}
	defer restore()
	h.logger.Infof("Successful uploaded file %v %v", file.URI, file.MIMEType)

	h.logger.Info("Generating AI response")
	result, err := h.GenAiClient.GenerateContent(file)
	if err != nil {
		h.logger.Errorf("Oops failed getting AI response due to %v", err)
		return
	}
	h.UpdateRichText(result.Text())
	h.logger.Info("Done generating AI response. Good Luck!")
}
