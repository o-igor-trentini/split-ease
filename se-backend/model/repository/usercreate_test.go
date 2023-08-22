package repository

import (
	"fmt"
	"se-backend/utils/utest"
	"testing"
)

func TestUserDomainRepository_Create(t *testing.T) {
	mock, db := utest.NewDBMock()
	fmt.Print(mock, db)
	//mock.exp
}
