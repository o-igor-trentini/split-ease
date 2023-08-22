package utest

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDBMock() (sqlmock.Sqlmock, *gorm.DB) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("[TESTE] Não foi possível gerar o mock do banco de dados [erro: %s]", err)
	}

	dialector := postgres.New(postgres.Config{Conn: mockDb, DriverName: "postgres"})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("[TESTE] Não foi possível abrir do banco de dados mockado [erro: %s]", err)
	}

	return mock, db
}
