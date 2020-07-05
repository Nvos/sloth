package infra

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level zapcore.Level

type Logger interface {
	SetLevel(Level)
	Background() *zap.Logger
}

type defaultLogger struct {
	logger *zap.Logger
	atom   *zap.AtomicLevel
}

func (l *defaultLogger) SetLevel(level Level) {
	l.atom.SetLevel(zapcore.Level(level))
}

func (l *defaultLogger) Background() *zap.Logger {
	return l.logger
}

func MustNewLogger(cfg *Configuration) Logger {
	var lvl zapcore.Level
	err := lvl.UnmarshalText([]byte(cfg.Logger.Level))
	if err != nil {
		panic(fmt.Errorf("failed to unmrashal logger level, err=%w", err))
	}

	logger, err := newLogger(cfg.Environment.Mode, lvl)
	if err != nil {
		panic(fmt.Errorf("failed to create new logger, err=%w", err))
	}

	return logger
}

func NewLogger(cfg *Configuration) (Logger, error) {
	var lvl zapcore.Level
	err := lvl.UnmarshalText([]byte(cfg.Logger.Level))
	if err != nil {
		return nil, fmt.Errorf("failed to unmrashal logger level, err=%w", err)
	}

	return newLogger(cfg.Environment.Mode, lvl)
}

func newLogger(mode Mode, lvl zapcore.Level) (Logger, error) {
	var loggerCfg zap.Config

	switch mode {
	case Develop, Test:
		loggerCfg = zap.NewDevelopmentConfig()
	case Production:
		loggerCfg = zap.NewProductionConfig()
	default:
		return nil, fmt.Errorf("unknown mode=%v", mode)
	}
	atom := zap.NewAtomicLevelAt(lvl)
	logger, err := loggerCfg.Build(zap.AddStacktrace(zap.DPanicLevel))
	if err != nil {
		return nil, err
	}

	return &defaultLogger{
		logger: logger,
		atom:   &atom,
	}, nil
}