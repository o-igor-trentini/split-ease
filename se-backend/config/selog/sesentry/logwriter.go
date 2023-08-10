package sesentry

import (
	"errors"
	"github.com/getsentry/sentry-go"
)

type LoggerOut struct {
}

func (t LoggerOut) Write(p []byte) (n int, err error) {
	var strP = string(p)
	var payload LoggerPayload
	payload.JSONToStruct(p)

	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(ZapLevelToSentryLevel[payload.Level])

		if payload.Tag != nil {
			scope.SetTag("tag", *payload.Tag)
		}

		sentry.CaptureException(errors.New(strP))
	})

	return len(strP), nil
}

func (t LoggerOut) Sync() error {
	return nil
}

func (t LoggerOut) Close() error {
	return nil
}
