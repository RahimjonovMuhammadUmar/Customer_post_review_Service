package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Field ...
type Field = zapcore.Field

var (
	// Int ...
	Int = zap.Int
	// String ...
	String = zap.String
	// Error ...
	Error = zap.Error
	// Bool ...
	Bool = zap.Bool

	// Any ...
	Any = zap.Any
)

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, field ...Field)
	Warn(msg string, field ...Field)
	Error(msg string, field ...Field)
	Fatal(msg string, field ...Field)
}

type LoggerImpl struct {
	zap *zap.Logger
}

var (
	customTimeFormat string
)

// New ...
func New(level, namespace string) Logger {
	if level == "" {
		level = LevelInfo
	}

	logger := LoggerImpl{
		zap: newZapLogger(level, time.RFC3339),
	}

	logger.zap = logger.zap.Named(namespace)

	zap.RedirectStdLog(logger.zap)

	return &logger
}

func (l *LoggerImpl) Debug(msg string, fields ...Field) {
	l.zap.Debug(msg, fields...)
}

func (l *LoggerImpl) Info(msg string, fields ...Field) {
	l.zap.Info(msg, fields...)
}

func (l *LoggerImpl) Warn(msg string, fields ...Field) {
	l.zap.Warn(msg, fields...)
}

func (l *LoggerImpl) Error(msg string, fields ...Field) {
	l.zap.Error(msg, fields...)
}

func (l *LoggerImpl) Fatal(msg string, fields ...Field) {
	l.zap.Fatal(msg, fields...)
}

// GetNamed ...
func GetNamed(l Logger, name string) Logger {
	switch v := l.(type) {
	case *LoggerImpl:
		v.zap = v.zap.Named(name)
		return v
	default:
		l.Info("logger GetNamed: invalid logger type")
		return l
	}
}

// WithFields ...
func WithFields(l Logger, fields ...Field) Logger {
	switch v := l.(type) {
	case *LoggerImpl:
		return &LoggerImpl{
			zap: v.zap.With(fields...),
		}
	default:
		l.Info("logger withFields: invalid logger type")
		return l
	}
}

// Cleanup
func Cleanup(l Logger) error {
	switch v := l.(type) {
	case *LoggerImpl:
		return v.zap.Sync()
	default:
		l.Info("logger cleanup invalid type logger")
		return nil
	}
}
