package utest

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RepositorySuite[R, M any] struct {
	suite.Suite
	conn  *sql.DB
	db    *gorm.DB
	Mock  sqlmock.Sqlmock
	Repo  R
	Model M
}

func NewRepositorySuite[R, M any](initRepo func(db *gorm.DB) R) *RepositorySuite[R, M] {
	var impl RepositorySuite[R, M]
	impl.setup(initRepo)

	return &impl
}

func (rs *RepositorySuite[R, M]) setup(initRepoFunc func(db *gorm.DB) R) {
	var err error

	rs.conn, rs.Mock, err = sqlmock.New()
	assert.NoError(rs.T(), err)

	dialector := postgres.New(postgres.Config{
		Conn:                 rs.conn,
		DriverName:           "postgres",
		DSN:                  "sqlmock_db_0",
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(rs.T(), err)

	rs.db = db
	rs.Repo = initRepoFunc(db)

	assert.NoError(rs.T(), err)
}

func (rs *RepositorySuite[R, M]) TearDownTest() {
	// Verifica se todas as exigências do teste foram compridas.
	assert.NoError(rs.T(), rs.Mock.ExpectationsWereMet())
}

func (rs *RepositorySuite[R, M]) TearDownSuite() {
	// Fecha a conexão com o banco de dados
	sqlDB, err := rs.db.DB()
	assert.NoError(rs.T(), err)

	err = sqlDB.Close()
	assert.NoError(rs.T(), err)
}
