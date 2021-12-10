package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func New(level string, path string) (*Logger, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		return nil, fmt.Errorf("cannot parse level: %w", err)
	}
	zerolog.SetGlobalLevel(lvl)
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(logFile)
	return &Logger{logger}, nil
}

func (l Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l Logger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

func (l Logger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}
