package screenshotter

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
	"path"
	"codenudge/pkg/config"
	"time"

	"github.com/kbinani/screenshot"
	"go.uber.org/zap"
)

type Screenshotter struct {
	Config *config.Screenshot
	logger *zap.SugaredLogger
}

func NewScreenshotter(cfg *config.Screenshot, logger *zap.SugaredLogger) *Screenshotter {
	return &Screenshotter{
		Config: cfg,
		logger: logger,
	}
}

func (sc *Screenshotter) Capture() (*bytes.Buffer, error) {
	bounds := screenshot.GetDisplayBounds(sc.Config.CaptureDisplayIndex)

	sc.logger.Info("Taking screenshot")
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	err = png.Encode(buf, img)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (sc *Screenshotter) SaveToDir(buf *bytes.Buffer) {
	fileName := fmt.Sprintf("%v.%v", time.Now().Format(sc.Config.TimeFormat), sc.Config.FileType)
	filePath := path.Join(sc.Config.SaveDir, fileName)

	sc.logger.Infof("Saving screenshot %v ", filePath)
	go func() {
		_, err := os.Stat(sc.Config.SaveDir)
		if !os.IsExist(err) {
			os.Mkdir(sc.Config.SaveDir, os.ModePerm)
		}

		savedF, err := os.Create(filePath)
		if err != nil {
			sc.logger.Errorf("Couldn't create file %v", err)
		}
		defer savedF.Close()

		_, err = savedF.Write(buf.Bytes())
		if err != nil {
			sc.logger.Errorf("Couldn't write screenshot to file %v", err)
		}
	}()
}

func (sc *Screenshotter) CaptureAndSafeInDir() (*bytes.Buffer, error) {
	buf, err := sc.Capture()
	if err != nil {
		return nil, err
	}
	if sc.Config.SaveScreenshots {
		sc.SaveToDir(buf)
	}

	return buf, err
}
