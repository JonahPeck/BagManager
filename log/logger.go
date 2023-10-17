package log

import (
	"golang.org/x/exp/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

// NewLogger will configure a new slog using the info level and text output mode
func NewLogger() *Logger {
	opts := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	sLogger := slog.New(opts.NewTextHandler(os.Stdout))
	return &Logger{logger: sLogger}
}

// Debug logs a message at level Debug on the logger
func (l *Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

// Warn logs a message at level Warn on the logger
func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

// Info logs a message at level Info on the logger
func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

// Error logs a message at level Error on the logger
func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

// Fatal logs an error message then calls os.Exit(1)
func (l *Logger) Fatal(msg string) {
	l.logger.Error(msg)
	os.Exit(1)
}

// logger.With("key1", value1, "key2", value2, ...)
func (l *Logger) With(args ...any) *Logger {
	var (
		attr  slog.Attr
		attrs []slog.Attr
	)
	for len(args) > 0 {
		attr, args = argsToSlogAttr(args)
		attrs = append(attrs, attr)
	}
	return &Logger{logger: slog.New(l.logger.Handler().WithAttrs(attrs))}
}

func (l *Logger) WithError(err error) *Logger {
	return l.With("error", err)
}

func (l *Logger) WithGroup(name string) *Logger {
	newLogger := l.logger.WithGroup(name)
	return &Logger{logger: newLogger}
}

func argsToSlogAttr(args []any) (slog.Attr, []any) {
	switch x := args[0].(type) {
	case string:
		if len(args) == 1 {
			return slog.String("!BADKEY", x), nil
		}
		a := slog.Any(x, args[1])
		a.Value = a.Value.Resolve()
		return a, args[2:]

	case slog.Attr:
		x.Value = x.Value.Resolve()
		return x, args[1:]

	default:
		return slog.Any("!BADKEY", x), args[1:]
	}
}
