package app

import (
	"context"
	"github.com/charmbracelet/log"
)

type Logger interface {
	Info(ctx context.Context, msg string)
	Debug(ctx context.Context, msg string)
	Warn(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)

	Infof(ctx context.Context, msg string, args ...any)
	Debugf(ctx context.Context, msg string, args ...any)
	Errorf(ctx context.Context, msg string, args ...any)
	Warnf(ctx context.Context, msg string, args ...any)
}

type DefaultLogger struct {
	logger *log.Logger
}

func NewDefaultLogger(prefix string) *DefaultLogger {
	logger := log.Default()
	logger.SetPrefix("【" + prefix + "】")
	return &DefaultLogger{
		logger: logger,
	}
}

func (l *DefaultLogger) Info(ctx context.Context, msg string) {
	l.logger.Info(msg)
}

func (l *DefaultLogger) Debug(ctx context.Context, msg string) {
	l.logger.Debug(msg)
}

func (l *DefaultLogger) Warn(ctx context.Context, msg string) {
	l.logger.Warn(ctx, msg)
}

func (l *DefaultLogger) Error(ctx context.Context, msg string) {
	l.logger.Error(ctx, msg)
}

func (l *DefaultLogger) Infof(ctx context.Context, msg string, args ...any) {
	l.logger.Infof(msg, args...)
}

func (l *DefaultLogger) Debugf(ctx context.Context, msg string, args ...any) {
	l.logger.Debugf(msg, args...)
}

func (l *DefaultLogger) Errorf(ctx context.Context, msg string, args ...any) {
	l.logger.Errorf(msg, args...)
}

func (l *DefaultLogger) Warnf(ctx context.Context, msg string, args ...any) {
	l.logger.Warnf(msg, args...)
}

var _ Logger = (*DefaultLogger)(nil)
