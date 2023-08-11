package sepostgresmigration

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"se-backend/config/seenv"
	"se-backend/config/selog"
	"se-backend/model/repository/entity"
	"time"
)

// CreateMigrationFile gera um arquivo ".go" pronto para escrever uma migration.
func CreateMigrationFile() {
	now := time.Now().Format("2006121545")
	fileName := fmt.Sprintf("migration%s", now)
	fileContent := []byte(fmt.Sprintf(
		`package migration

import "gorm.io/gorm"

func %s() MigrationExecute {
	return MigrationExecute{
		Name: "%s",
		Run: func(db *gorm.DB) error {
			return nil
		},
	}
}
	`,
		fileName,
		fileName,
	))

	if err := os.WriteFile("./config/database/sepostgres/sepostgresmigration/migrations/"+fileName+".go", fileContent, 0644); err != nil {
		selog.Fatal("Não foi possível criar o arquivo de migração", err, selog.MigrationTag)
	}
}

// CreateDatabaseStructure cria a estrutura do banco.
func CreateDatabaseStructure(db *gorm.DB) {
	handleStepError := func(tx *gorm.DB, message string, err error) {
		if err != nil {
			tx.Rollback()
			selog.Fatal(message, err)
		}
	}

	tx := db.Begin()

	handleStepError(tx, "Não foi possível criar o schema principal", createMainSchema(tx))
	handleStepError(tx, "Não foi possível criar os enumerators", createEnums(tx))
	handleStepError(tx, "Não foi possível criar as tabelas", createTables(tx))

	tx.Commit()
}

// createMainSchema cria o esquema principal caso não exista.
func createMainSchema(db *gorm.DB) error {
	return db.Exec("CREATE SCHEMA IF NOT EXISTS " + seenv.ENV.DatabaseSchema).Error
}

// createEnums cria os enumerators caso não existam.
func createEnums(db *gorm.DB) error {
	return nil
}

// createTables cria as tabelas caso não existam.
// Sempre que um model (referente a uma tabela) for criado, ele deve ser referênciado aqui.
func createTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.Migration{},
		&entity.UserEntity{},
	)
}
