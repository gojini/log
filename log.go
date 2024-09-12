package sbomd

import (
	"context"
	"io"
	"log/slog"
	"os"
)

type LogKeyType string

const LogKey = LogKeyType("slog")

func Log(ctx context.Context) *slog.Logger {
	log := ctx.Value(LogKey)
	if log != nil {
		logger, ok := log.(*slog.Logger)

		if !ok {
			return StdoutLogger()
		}

		return logger
	}

	return StdoutLogger()
}

func NilLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(io.Discard, nil))
}

func StdoutLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
