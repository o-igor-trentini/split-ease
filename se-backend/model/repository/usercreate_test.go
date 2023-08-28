package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"se-backend/model"
	"se-backend/model/repository/entity/entityconverter"
	"se-backend/utils/utest"
	"testing"
	"time"
)

//func TestUserDomainRepository_Create(t *testing.T) {
//	t.Run("Insere com sucesso", func(t *testing.T) {
//		var rs repositorySuite
//		rs.Setup()
//
//		user := model.NewUserDomain(
//			"Igor",
//			"Trentini",
//			"igor@gmail.com",
//			"igor",
//			"Senha@123",
//		)
//		user.EncryptPassword()
//
//		now := time.Now()
//
//		rs.mock.ExpectBegin()
//		rs.mock.
//			ExpectExec(regexp.QuoteMeta(`INSERT INTO "us_users" ("us_created_at","us_updated_at","us_deleted_at","us_first_name","us_last_name","us_email","us_username","us_password") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "us_id"`)).
//			WithArgs(
//				now,
//				now,
//				nil,
//				user.GetFirstName(),
//				user.GetLastName(),
//				user.GetEmail(),
//				user.GetUsername(),
//				user.GetPassword(),
//			).
//			WillReturnResult(sqlmock.NewResult(1, 1))
//		rs.mock.ExpectCommit()
//
//		r := NewUserDomain(rs.DB)
//
//		userEntity, err := r.Create(user)
//		assert.NoError(s.T(), err) // avalia se não houve nenhum erro na execução
//		//assert.Equal(rs.T(), rs.person, p) // verificar se as struts são iguais
//
//		fmt.Print(userEntity)
//	})
//}

func TestUserDomainRepository_FindOneByEmail(t *testing.T) {
	t.Run("Busca com sucesso", func(t *testing.T) {
		s := utest.NewRepositorySuite[UserDomainRepository, model.UserDomainInterface](NewUserDomain)

		input := "trentini@gmail.com"

		userDomain := model.NewUserDomain(
			"Igor",
			"Trentini",
			"igor@gmail.com",
			"igor",
			"Senha@123",
		)
		userDomain.SetID(1)
		userDomain.EncryptPassword()

		now := time.Now()

		rows := sqlmock.
			NewRows([]string{
				"us_id",
				"us_created_at",
				"us_updated_at",
				"us_deleted_at",
				"us_first_name",
				"us_last_name",
				"us_email",
				"us_username",
				"us_password",
			}).
			AddRow(
				userDomain.GetID(),
				now,
				now,
				nil,
				userDomain.GetFirstName(),
				userDomain.GetLastName(),
				userDomain.GetEmail(),
				userDomain.GetUsername(),
				userDomain.GetPassword(),
			)
		s.Mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "us_users" WHERE "us_email" = $1 AND "us_users"."us_deleted_at" IS NULL ORDER BY "us_users"."us_id" LIMIT 1`)).
			WithArgs(input).
			WillReturnRows(rows)

		// TODO: Estou colocando um input diferente do valor e está retornando o usuário com o valor
		// Corrigir isso

		userEntity, err := s.Repo.FindOneByEmail(input)
		result := entityconverter.UserEntityToDomain(userEntity)

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), userDomain, result)
	})
}
