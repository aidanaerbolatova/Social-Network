package logger

import (
	"context"

	"go.uber.org/zap"
)

type correlationIdType int

const (
	requestIdKey correlationIdType = iota
	sessionIdKey
)

var logger zap.SugaredLogger

func init() {
	zapLogger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	logger = *zapLogger.Sugar()
}

// WithRqId returns a context which knows its request ID
func WithRqId(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, requestIdKey, reqID)
}

// WithSessionId returns a context which knows its session ID
func WithSessionId(ctx context.Context, sessionId string) context.Context {
	return context.WithValue(ctx, sessionIdKey, sessionId)
}

// Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) zap.SugaredLogger {
	newLogger := &logger
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(requestIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("rqId", ctxRqId))
		}
		if ctxSessionId, ok := ctx.Value(sessionIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("sessionId", ctxSessionId))
		}
	}
	return *newLogger
}
