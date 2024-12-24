package app

type Option func(app *App) error

func WithLogger(logger Logger) Option {
	return func(app *App) error {
		app.logger = logger
		return nil
	}
}
