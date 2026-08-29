package slog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
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

func (l Manager) Infof(format string, args ...any) {
	l.logger.InfoContext(context.Background(), l.addCallerInfo(fmt.Sprintf(format, args...)))
}

func (l Manager) Errorf(format string, args ...any) {
	l.logger.ErrorContext(context.Background(), l.addCallerInfo(fmt.Sprintf(format, args...)))
}

func (l Manager) Fatalf(format string, args ...any) {
	l.logger.ErrorContext(context.Background(), l.addCallerInfo(fmt.Sprintf(format, args...)))
	os.Exit(1)
}

func (l Manager) Debugf(format string, args ...any) {
	l.logger.DebugContext(context.Background(), l.addCallerInfo(fmt.Sprintf(format, args...)))
}

func (l Manager) Info(args ...any) {
	l.logger.InfoContext(context.Background(), l.addCallerInfo(fmt.Sprint(args...)))
}

func (l Manager) Error(args ...any) {
	l.logger.ErrorContext(context.Background(), l.addCallerInfo(fmt.Sprint(args...)))
}

func (l Manager) Fatal(args ...any) {
	l.logger.ErrorContext(context.Background(), l.addCallerInfo(fmt.Sprint(args...)))
}

func (l Manager) Debug(args ...any) {
	l.logger.DebugContext(context.Background(), l.addCallerInfo(fmt.Sprint(args...)))
}

func (l Manager) addCallerInfo(msg string) string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return msg
	}
	file = filepath.Base(file)
	return fmt.Sprintf("%s:%d %s", file, line, msg)
}
