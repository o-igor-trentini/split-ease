package selog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/url"
	"se-backend/config/seenv"
	"se-backend/config/selog/sesentry"
)

var logger *zap.Logger

func Init() {
	cfg := zap.Config{
		OutputPaths: getLogOutputs(),
		Level:       zap.NewAtomicLevelAt(getLogLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:      "level",
			TimeKey:       "time",
			MessageKey:    "message",
			StacktraceKey: "stackTrace",
			FunctionKey:   "function",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	applyDevelopmentConfig(&cfg)

	logger = buildLogger(cfg)

	if err := sesentry.Init(); err != nil {
		Error("Não foi possível inicializar o Sentry", err)
	}
}

// getLogOutputs informa as saídas do logger.
func getLogOutputs() []string {
	customWriterKey := "sentry"
	customWriterPath := customWriterKey + ":custom"
	registerCustomWriterKey(customWriterKey)

	return []string{seenv.ENV.LogOutPut, customWriterPath}
}

// getLogLevel informa o nível do logger.
func getLogLevel() zapcore.Level {
	var LogLevelToZapLevels = map[string]zapcore.Level{
		seenv.LogLevelDebug: zapcore.DebugLevel,
		seenv.LogLevelInfo:  zapcore.InfoLevel,
		seenv.LogLevelError: zapcore.ErrorLevel,
		seenv.LogLevelFatal: zapcore.ErrorLevel,
	}

	return LogLevelToZapLevels[seenv.ENV.LogLevel]
}

// applyDevelopmentConfig aplica as particularidades do ambiente de desenvolvimento.
func applyDevelopmentConfig(cfg *zap.Config) {
	if seenv.ENV.AppEnvironment == seenv.AppEnvironmentProduction {
		return
	}

	cfg.Development = true
	cfg.DisableStacktrace = false
}

// registerCustomWriterKey registra uma chave personalizada para outra saída de log.
func registerCustomWriterKey(customWriterKey string) {
	if err := zap.RegisterSink(customWriterKey, func(u *url.URL) (zap.Sink, error) { return sesentry.LoggerOut{}, nil }); err != nil {
		log.Fatalf("erro ao registrar chave personalizada de log [erro: %s]", err)
	}
}

// buildLogger gera o logger.
func buildLogger(cfg zap.Config) *zap.Logger {
	zapLogger, err := cfg.Build()
	if err != nil {
		log.Fatalf("erro ao inicializar o serviço de logging [erro: %s]", err)
	}

	return zapLogger
}

func loggerIsValid() bool {
	return logger != nil
}

func Info(message string, tags ...zap.Field) {
	if !loggerIsValid() {
		return
	}

	logger.Info(message, tags...)
	_ = logger.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	if !loggerIsValid() {
		return
	}

	tags = append(tags, zap.NamedError("error", err))
	logger.Error(message, tags...)
	_ = logger.Sync()
}

func Fatal(message string, err error, tags ...zap.Field) {
	if !loggerIsValid() {
		return
	}

	tags = append(tags, zap.NamedError("fatal", err))
	logger.Fatal(message, tags...)
	_ = logger.Sync()
}
