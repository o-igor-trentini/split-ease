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

	// Banco de dados
	envDatabaseHost     string = "DB_HOST"
	envDatabasePort     string = "DB_PORT"
	envDatabaseUser     string = "DB_USER"
	envDatabasePassword string = "DB_PASSWORD"
	envDatabaseName     string = "DB_NAME"
	envDatabaseSchema   string = "DB_SCHEMA"

	// JWT
	envJWTSecretKey string = "JWT_SECRET_KEY"

	// Email
	envEmailAPIKey string = "EMAIL_API_KEY"
	envEmailFrom   string = "EMAIL_FROM"

	// Frontend
	envFrontendURLKey string = "FRONTEND_URL"
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
	DatabaseHost      string
	DatabasePort      string
	DatabaseUser      string
	DatabasePassword  string
	DatabaseName      string
	DatabaseSchema    string
	JWTSecretKey      string
	EmailAPIKey       string
	EmailFrom         string
	FrontendURL       string
}
