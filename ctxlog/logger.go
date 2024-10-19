package ctxlog

import (
	"context"
	"github.com/pr0n1x/go-liners/logr"
)

type ContextKey string

const CtxKeyLogger ContextKey = "ctxlog.logger"

type loggerCtxWrapper struct {
	l logr.Logger
}

func CtxLogger(ctx context.Context) logr.Logger {
	if v, ok := ctx.Value(CtxKeyLogger).(loggerCtxWrapper); ok {
		return v.l
	}
	return logr.ZeroLogger{}
}

func CtxNonZeroLogger(ctx context.Context) (bool, logr.Logger) {
	if v, ok := ctx.Value(CtxKeyLogger).(loggerCtxWrapper); ok {
		return true, v.l
	}
	return false, logr.ZeroLogger{}
}

func WithLogger(ctx context.Context, logger logr.Logger) context.Context {
	return context.WithValue(ctx, CtxKeyLogger, loggerCtxWrapper{l: logger})
}
