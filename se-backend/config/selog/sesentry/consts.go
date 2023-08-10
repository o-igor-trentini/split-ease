package sesentry

import (
	"github.com/getsentry/sentry-go"
	"se-backend/config/seenv"
)

var ZapLevelToSentryLevel = map[string]sentry.Level{
	seenv.LogLevelDebug: sentry.LevelDebug,
	seenv.LogLevelInfo:  sentry.LevelInfo,
	seenv.LogLevelError: sentry.LevelError,
	seenv.LogLevelFatal: sentry.LevelFatal,
}
