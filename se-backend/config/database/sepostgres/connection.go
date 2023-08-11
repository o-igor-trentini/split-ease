package sepostgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"se-backend/config/seenv"
	"se-backend/config/selog"
	"time"
)

func New() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s search_path=%s",
		seenv.ENV.DatabaseHost,
		seenv.ENV.DatabaseUser,
		seenv.ENV.DatabasePassword,
		seenv.ENV.DatabaseName,
		seenv.ENV.DatabasePort,
		seenv.ENV.DatabaseSchema,
	)

	db, err := gorm.Open(postgres.Open(dsn), getDBCfg())
	if err != nil {
		selog.Fatal("Não foi possível conectar ao banco de dados", err)
	}

	return db
}

// getDBCfg gera a configuração do banco de dados.
func getDBCfg() *gorm.Config {
	var cfg gorm.Config

	if seenv.ENV.AppEnvironment == seenv.AppEnvironmentProduction {
		return &cfg
	}

	var levelsMapping = map[string]logger.LogLevel{
		seenv.LogLevelDebug: logger.Info,
		seenv.LogLevelInfo:  logger.Info,
		seenv.LogLevelError: logger.Error,
		seenv.LogLevelFatal: logger.Silent,
	}

	cfg.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,                       // Slow SQL threshold
			LogLevel:                  levelsMapping[seenv.ENV.LogLevel], // Log level
			IgnoreRecordNotFoundError: true,                              // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                              // Disable color
		},
	)

	return &cfg
}
