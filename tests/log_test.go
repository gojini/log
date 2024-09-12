package logger_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
	log "gojini.dev/log"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	require := require.New(t)

	logger := log.Log(ctx)
	require.NotNil(logger)
	require.Equal(logger, log.StdoutLogger())

	ctx = context.WithValue(ctx, log.LogKey, slog.Default())
	logger = log.Log(ctx)
	require.NotNil(logger)
	require.NotEqual(logger, log.StdoutLogger())
	require.Equal(logger, slog.Default())

	ctx = context.WithValue(ctx, log.LogKey, "hello")
	require.Equal(log.StdoutLogger(), log.Log(ctx))
}

func TestNilloggerger(t *testing.T) {
	t.Parallel()

	nilloggerger := log.NilLogger()
	require.NotNil(t, nilloggerger)
}

func TestStdOutloggerger(t *testing.T) {
	t.Parallel()

	require := require.New(t)
	require.NotNil(log.StdoutLogger())
}
