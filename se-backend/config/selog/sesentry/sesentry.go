package sesentry

import (
	"github.com/getsentry/sentry-go"
	"se-backend/config/seenv"
	"time"
)

// Init inicializa o logger do sentry.
func Init() error {
	sentrySyncTransport := sentry.NewHTTPSyncTransport()
	sentrySyncTransport.Timeout = time.Second * 3

	return sentry.Init(sentry.ClientOptions{
		Dsn:         seenv.ENV.SentryDSN,
		Environment: seenv.ENV.SentryEnvironment,
		Transport:   sentrySyncTransport,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
}
