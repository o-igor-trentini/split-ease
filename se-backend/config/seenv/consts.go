package seenv

// Nome das variáveis de ambiente
const (
	// Aplicação
	envAppEnvironment string = "APP_ENVIRONMENT"

	// Servidor
	envServerMode string = "SERVER_MODE"
	envServerHost string = "SERVER_HOST"
	envServerPort string = "SERVER_PORT"

	// Log
	envLogLevel  string = "LOG_LEVEL"
	envLogOutput string = "LOG_OUTPUT"

	// Sentry
	envSentryDSN         string = "SENTRY_DSN"
	envSentryEnvironment string = "SENTRY_ENVIRONMENT"
)

// Possíveis valores das variáveis de ambiente
const (
	AppEnvironmentDevelopment string = "development"
	AppEnvironmentStaging     string = "staging"
	AppEnvironmentProduction  string = "production"

	ServerModeDebug   string = "debug"
	ServerModeRelease string = "release"

	LogLevelDebug string = "debug"
	LogLevelInfo  string = "info"
	LogLevelError string = "error"
	LogLevelFatal string = "fatal"
	// LogLevelWarn   string = "warn"
	//LogLevelPanic  string = "panic"
	//LogLevelDPanic string = "dpanic"

	LogOutputStdout string = "stdout"

	SentryEnvironmentDevelopment string = "development"
	SentryEnvironmentStaging     string = "staging"
	SentryEnvironmentProduction  string = "production"
)

type ENVType struct {
	AppEnvironment    string
	ServerMode        string
	ServerHost        string
	ServerPort        string
	LogLevel          string
	LogOutPut         string
	SentryDSN         string
	SentryEnvironment string
}
