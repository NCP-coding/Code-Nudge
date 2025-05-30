package shortcut

import (
	"context"
	"reflect"
	"codenudge/pkg/config"
	"codenudge/pkg/handler"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.design/x/hotkey"
)

type Shortcut struct {
	Config      *config.Shortcut
	Handler     *handler.Handler
	SelectCases []reflect.SelectCase
	logger      *zap.SugaredLogger
}

func NewShortcut(cfg *config.Shortcut, handler *handler.Handler, logger *zap.SugaredLogger) *Shortcut {
	return &Shortcut{
		Config:  cfg,
		Handler: handler,
		logger:  logger,
	}
}

func (sc *Shortcut) Register() {
	sc.SelectCases = make([]reflect.SelectCase, len(sc.Config.HotkeyConfigs))
	for i, cfg := range sc.Config.HotkeyConfigs {
		key := hotkey.New(cfg.KeyModifier, cfg.Key)
		err := key.Register()
		if err != nil {
			sc.logger.Error("Cannot register hotkey %v", cfg.Name)
			continue
		}
		sc.SelectCases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(key.Keydown()),
		}
	}
}
func Run(lc fx.Lifecycle, logger *zap.SugaredLogger, shortcut *Shortcut) error {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Infoln("Started event loop for shortcuts")
			go func() {
				for {
					chosen, _, ok := reflect.Select(shortcut.SelectCases)
					if !ok {
						continue
					}
					cfg := shortcut.Config.HotkeyConfigs[chosen]
					logger.Infof("Recieved event for %v", cfg.Name)
					fn, ok := shortcut.Handler.Callbacks[cfg.Name]
					if ok {
						fn()
					} else {
						logger.Fatalf("Hhandler with name %v does not exist, check config", cfg.Name)
					}

				}
			}()
			return nil
		},
	})
	return nil
}
