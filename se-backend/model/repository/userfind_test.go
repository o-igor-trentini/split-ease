package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"regexp"
	"se-backend/config/seerror"
	"se-backend/model"
	"se-backend/model/repository/entity/entityconverter"
	"se-backend/utils/utest"
	"testing"
	"time"
)

func TestUserDomainRepository_FindOneByUsername(t *testing.T) {
	type args struct {
		username string
	}

	type test struct {
		name      string
		want      model.UserDomainInterface
		expectErr utest.Error
		args      args
	}

	userDomain := model.NewUserDomain(
		"Igor",
		"Trentini",
		"igor@gmail.com",
		"igor",
		"Senha@123",
	)
	userDomain.SetID(1)
	userDomain.EncryptPassword()

	table := []test{
		{
			name: "Sucesso informando nome de usuário correto",
			want: userDomain,
			args: args{
				username: "igor",
			},
		},
		{
			name: "Falha informando nome de usuário não existente",
			want: nil,
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewNotFoundErr("Usuário não encontrado", gorm.ErrRecordNotFound),
				MockValue: gorm.ErrRecordNotFound,
			},
			args: args{
				username: "trentini",
			},
		},
		{
			name: "Falha por instância do banco de dados inválida",
			want: nil,
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewsInternalServerErrorErr("Não foi possível buscar o usuário", gorm.ErrInvalidDB),
				MockValue: gorm.ErrInvalidDB,
			},
		},
		{
			name: "Falha por transação inválida",
			want: nil,
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewsInternalServerErrorErr("Não foi possível buscar o usuário", gorm.ErrInvalidTransaction),
				MockValue: gorm.ErrInvalidTransaction,
			},
		},
	}

	s := utest.NewRepositorySuite[UserDomainRepository, model.UserDomainInterface](t, NewUserDomain)

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			var rows = new(sqlmock.Rows)
			if test.want != nil {
				now := time.Now()

				rows = sqlmock.
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
						test.want.GetID(),
						now,
						now,
						nil,
						test.want.GetFirstName(),
						test.want.GetLastName(),
						test.want.GetEmail(),
						test.want.GetUsername(),
						test.want.GetPassword(),
					)
			}

			s.Mock.
				ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "us_users" WHERE "us_username" = $1 AND "us_users"."us_deleted_at" IS NULL ORDER BY "us_users"."us_id" LIMIT 1`),
				).
				WithArgs(test.args.username).
				WillReturnError(test.expectErr.MockValue).
				WillReturnRows(rows)

			userEntity, err := s.Repo.FindOneByUsername(test.args.username)
			result := entityconverter.UserEntityToDomain(userEntity)

			if test.expectErr.Expected {
				assert.Zero(s.T(), userEntity)
				assert.Equal(s.T(), test.expectErr.Value, err)
			} else {
				assert.NoError(s.T(), err)
				assert.Equal(s.T(), test.want, result)
			}
		})
	}
}

func TestUserDomainRepository_FindOneByEmail(t *testing.T) {
	type args struct {
		email string
	}

	type test struct {
		name      string
		want      model.UserDomainInterface
		expectErr utest.Error
		args      args
	}

	userDomain := model.NewUserDomain(
		"Igor",
		"Trentini",
		"igor@gmail.com",
		"igor",
		"Senha@123",
	)
	userDomain.SetID(1)
	userDomain.EncryptPassword()

	table := []test{
		{
			name: "Sucesso informando e-mail correto",
			want: userDomain,
			args: args{
				email: "igor@gmail.com",
			},
		},
		{
			name: "Falha informando e-mail não existente",
			want: nil,
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewNotFoundErr("Usuário não encontrado", gorm.ErrRecordNotFound),
				MockValue: gorm.ErrRecordNotFound,
			},
			args: args{
				email: "trentini@gmail.com",
			},
		},
		{
			name: "Falha por instância do banco de dados inválida",
			want: nil,
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewsInternalServerErrorErr("Não foi possível buscar o usuário", gorm.ErrInvalidDB),
				MockValue: gorm.ErrInvalidDB,
			},
		},
		{
			name: "Falha por transação inválida",
			want: nil,
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewsInternalServerErrorErr("Não foi possível buscar o usuário", gorm.ErrInvalidTransaction),
				MockValue: gorm.ErrInvalidTransaction,
			},
		},
	}

	s := utest.NewRepositorySuite[UserDomainRepository, model.UserDomainInterface](t, NewUserDomain)

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			var rows = new(sqlmock.Rows)
			if test.want != nil {
				now := time.Now()

				rows = sqlmock.
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
						test.want.GetID(),
						now,
						now,
						nil,
						test.want.GetFirstName(),
						test.want.GetLastName(),
						test.want.GetEmail(),
						test.want.GetUsername(),
						test.want.GetPassword(),
					)
			}

			s.Mock.
				ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "us_users" WHERE "us_email" = $1 AND "us_users"."us_deleted_at" IS NULL ORDER BY "us_users"."us_id" LIMIT 1`),
				).
				WithArgs(test.args.email).
				WillReturnError(test.expectErr.MockValue).
				WillReturnRows(rows)

			userEntity, err := s.Repo.FindOneByEmail(test.args.email)
			result := entityconverter.UserEntityToDomain(userEntity)

			if test.expectErr.Expected {
				assert.Zero(s.T(), userEntity)
				assert.Equal(s.T(), test.expectErr.Value, err)
			} else {
				assert.NoError(s.T(), err)
				assert.Equal(s.T(), test.want, result)
			}
		})
	}
}
