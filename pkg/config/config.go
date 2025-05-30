package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"golang.design/x/hotkey"
)

type HotkeyConfig struct {
	Name        string
	KeyModifier []hotkey.Modifier
	Key         hotkey.Key
	Callback    func()
}

type Shortcut struct {
	HotkeyConfigs []HotkeyConfig
}

type GenaiClient struct {
	ApiKey       string `env:"API_KEY,required"`
	Model        string `env:"MODEL,required"`
	ModelPrompt  string `env:"MODEL_PROMPT,required"`
	FileMimeType string
}

type Screenshot struct {
	SaveDir             string `env:"DIR,required"`
	CaptureDisplayIndex int    `env:"CAPTURE_DISPLAY_INDEX,required"`
	SaveScreenshots     bool   `env:"SAVE_ENABLED,required"`
	TimeFormat          string
	FileType            string
}

type MainWindow struct {
	WindowWidth  float32
	WindowHeight float32
	Title        string `env:"TITLE,required"`
}

type FyneApp struct {
	Id string `env:"ID,required"`
}

type Config struct {
	FyneApp     *FyneApp    `envPrefix:"FYNE_APP_"`
	MainWindow  *MainWindow `envPrefix:"MAIN_WINDOW_"`
	Shortcut    *Shortcut
	GenaiClient *GenaiClient `envPrefix:"GENAI_"`
	Screenshot  *Screenshot  `envPrefix:"SCREENSHOT_"`
	Logger      *zap.Logger
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Cannot load .env file %v", err)
	}

	cfg := &Config{
		FyneApp: &FyneApp{},
		MainWindow: &MainWindow{
			WindowWidth:  600,
			WindowHeight: 400,
		},
		Shortcut: &Shortcut{
			HotkeyConfigs: []HotkeyConfig{
				{
					Name:        "Alt + T: Example HINT",
					KeyModifier: []hotkey.Modifier{hotkey.ModAlt},
					Key:         hotkey.KeyT,
				},
				{
					Name:        "Ctrl + Alt + A: Screenshot",
					KeyModifier: []hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt},
					Key:         hotkey.KeyA,
				},
			},
		},
		GenaiClient: &GenaiClient{
			FileMimeType: "image/png",
		},
		Screenshot: &Screenshot{
			TimeFormat: "2006-01-02_15-04-05",
			FileType:   "png",
		},
	}

	err = env.Parse(cfg)
	if err != nil {
		log.Fatalf("Cannot parse env vars %v", err)
	}

	return cfg

}
