package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/dida-social/dida-server/pkg/frame"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// App application
type App struct {
	name    string
	logger  Logger
	modules []frame.Module
}

func (a *App) RegisterModules(ctx context.Context, modules ...frame.Module) error {
	a.modules = append(a.modules, modules...)
	return nil
}

func (a *App) RegisterReleaseFunc(releaseFunc ...frame.ReleaseFunc) error {

	return nil
}

func (a *App) Name() string {
	return a.name
}

func (a *App) Run(ctx context.Context) error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT /*ctrl-c*/, syscall.SIGQUIT, /*ctrl-/*/
		syscall.SIGKILL, syscall.SIGTERM)

	// TODO: start modules
	for _, module := range a.modules {
		if err := module.Run(ctx); err != nil {
			return err
		}
		a.logger.Infof(ctx, "%s is starting", module.Name())
	}

	a.logger.Info(ctx, "server is starting...")

	for {
		select {
		case sig := <-ch:
			switch sig {
			case syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM:
				// TODO: release resources
				var sb strings.Builder
				for _, module := range a.modules {
					if err := module.Close(); err != nil {
						a.logger.Errorf(ctx, "close module %s failed with error: %v", module.Name(), err)
						sb.WriteString(fmt.Sprintf("module %s close failed", module.Name()))
					}
				}
				if sb.Len() > 0 {
					return errors.New(sb.String())
				}
				return nil
			case syscall.SIGQUIT:
				// start online profiling
				continue
			}
		}
	}
}

var _ frame.AppModule = (*App)(nil)

func New(name string, opts ...Option) (*App, error) {
	logger := NewDefaultLogger(name)
	app := &App{
		name:   name,
		logger: logger,
	}

	for _, opt := range opts {
		if err := opt(app); err != nil {
			return nil, err
		}
	}

	return app, nil
}
