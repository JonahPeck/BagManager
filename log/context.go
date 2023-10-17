package log

import (
	"context"
	"os"

	"golang.org/x/exp/slog"
)

// LogIntoCtx returns a new context with the given logger.
func LogIntoCtx(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, "logger", logger)
}

// LogFromCtx returns a Logger from a context. If no logger is found in the context,
// a new one is created using a default text handler.
func LogFromCtx(ctx context.Context) *Logger {
	v, ok := ctx.Value("logger").(*Logger)
	if ok {
		return v
	}
	return &Logger{logger: slog.New(slog.NewTextHandler(os.Stdout))}
}
