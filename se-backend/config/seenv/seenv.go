package seenv

import (
	"github.com/joho/godotenv"
	"log"
	"se-backend/utils/uenv"
)

var ENV = new(ENVType)

// Init inicializa o arquivo de variáveis de ambiente.
func Init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("[ENV] Não foi possível carregar o arquivo de variáveis de ambiente [erro: %s]", err)
	}

	list := []uenv.CheckEnvTable{
		// Aplicação
		{
			Name:     envAppEnvironment,
			Expected: []string{AppEnvironmentDevelopment, AppEnvironmentStaging, AppEnvironmentProduction},
		},

		// Servidor
		{
			Name:     envServerMode,
			Expected: []string{ServerModeDebug, ServerModeRelease},
		},
		{Name: envServerHost, IgnoreEmpty: true},
		{Name: envServerPort, IgnoreEmpty: true},

		// Log
		{
			Name:     envLogLevel,
			Expected: []string{LogLevelDebug, LogLevelInfo, LogLevelError, LogLevelFatal},
		},
		{Name: envLogOutput, Expected: []string{LogOutputStdout}},

		// Sentry
		{Name: envSentryDSN},
		{
			Name:     envSentryEnvironment,
			Expected: []string{SentryEnvironmentDevelopment, SentryEnvironmentStaging, SentryEnvironmentProduction},
		},

		// Banco de dados
		{Name: envDatabaseHost},
		{Name: envDatabasePort},
		{Name: envDatabaseUser},
		{Name: envDatabasePassword},
		{Name: envDatabaseName},
		{Name: envDatabaseSchema},

		// JWT
		{Name: envJWTSecretKey},

		// Email
		{Name: envEmailAPIKey},
		{Name: envEmailFrom},

		// Frontend
		{Name: envFrontendURLKey},
	}

	values, err := uenv.CheckEnvs(list)
	if err != nil {
		log.Fatalf("[ENV] Não foi possível validar as variáveis de ambiente [erro: %s]", err)
	}

	ENV = &ENVType{
		AppEnvironment:    values[envAppEnvironment],
		ServerMode:        values[envServerMode],
		ServerHost:        values[envServerHost],
		ServerPort:        values[envServerPort],
		LogLevel:          values[envLogLevel],
		LogOutPut:         values[envLogOutput],
		SentryDSN:         values[envSentryDSN],
		SentryEnvironment: values[envSentryEnvironment],
		DatabaseHost:      values[envDatabaseHost],
		DatabasePort:      values[envDatabasePort],
		DatabaseUser:      values[envDatabaseUser],
		DatabasePassword:  values[envDatabasePassword],
		DatabaseName:      values[envDatabaseName],
		DatabaseSchema:    values[envDatabaseSchema],
		JWTSecretKey:      values[envJWTSecretKey],
		EmailAPIKey:       values[envEmailAPIKey],
		EmailFrom:         values[envEmailFrom],
		FrontendURL:       values[envFrontendURLKey],
	}
}
