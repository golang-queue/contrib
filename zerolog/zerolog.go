package zerolog

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// New to create new interface for logger
func New() *Manager {
	return &Manager{}
}

// Manager for zerolog
type Manager struct{}

func (l Manager) Infof(format string, args ...interface{}) {
	log.Info().Caller().Msgf(format, args...)
}

func (l Manager) Errorf(format string, args ...interface{}) {
	log.Error().Caller().Msgf(format, args...)
}

func (l Manager) Fatalf(format string, args ...interface{}) {
	log.Fatal().Caller().Msgf(format, args...)
}

func (l Manager) Debugf(format string, args ...interface{}) {
	log.Debug().Caller().Msgf(format, args...)
}

func (l Manager) Info(args ...interface{}) {
	log.Info().Caller().Msg(fmt.Sprint(args...))
}

func (l Manager) Error(args ...interface{}) {
	log.Error().Caller().Msg(fmt.Sprint(args...))
}

func (l Manager) Fatal(args ...interface{}) {
	log.Fatal().Caller().Msg(fmt.Sprint(args...))
}

func (l Manager) Debug(args ...interface{}) {
	log.Debug().Caller().Msg(fmt.Sprint(args...))
}
