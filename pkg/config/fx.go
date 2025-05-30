package config

import "go.uber.org/fx"

func FxModule() fx.Option {
	return fx.Module("configs",
		fx.Provide(
			Load,
			func(cfg *Config) *FyneApp { return cfg.FyneApp },
			func(cfg *Config) *MainWindow { return cfg.MainWindow },
			func(cfg *Config) *Shortcut { return cfg.Shortcut },
			func(cfg *Config) *GenaiClient { return cfg.GenaiClient },
			func(cfg *Config) *Screenshot { return cfg.Screenshot },
		),
	)
}
