package slog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// New to create new interface for logger
func New() *Manager {
	return &Manager{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

// Manager for slog
type Manager struct {
	logger *slog.Logger
}

func (l Manager) Infof(format string, args ...interface{}) {
	l.logger.InfoContext(context.Background(), fmt.Sprintf(format, args...))
}

func (l Manager) Errorf(format string, args ...interface{}) {
	l.logger.ErrorContext(context.Background(), fmt.Sprintf(format, args...))
}

func (l Manager) Fatalf(format string, args ...interface{}) {
	l.logger.ErrorContext(context.Background(), fmt.Sprintf(format, args...))
	os.Exit(1)
}

func (l Manager) Debugf(format string, args ...interface{}) {
	l.logger.DebugContext(context.Background(), fmt.Sprintf(format, args...))
}

func (l Manager) Info(args ...interface{}) {
	l.logger.InfoContext(context.Background(), fmt.Sprint(args...))
}

func (l Manager) Error(args ...interface{}) {
	l.logger.ErrorContext(context.Background(), fmt.Sprint(args...))
}

func (l Manager) Fatal(args ...interface{}) {
	l.logger.ErrorContext(context.Background(), fmt.Sprint(args...))
	os.Exit(1)
}

func (l Manager) Debug(args ...interface{}) {
	l.logger.DebugContext(context.Background(), fmt.Sprint(args...))
}
