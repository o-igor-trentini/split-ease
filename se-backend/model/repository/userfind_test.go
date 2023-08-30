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
			args: args{username: "igor"},
		},
		{
			name: "Falha informando nome de usuário não existente",
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewNotFoundErr("Usuário não encontrado", gorm.ErrRecordNotFound),
				MockValue: gorm.ErrRecordNotFound,
			},
			args: args{username: "trentini"},
		},
		{
			name: "Falha por instância do banco de dados inválida",
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewsInternalServerErrorErr("Não foi possível buscar o usuário", gorm.ErrInvalidDB),
				MockValue: gorm.ErrInvalidDB,
			},
		},
		{
			name: "Falha por transação inválida",
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewsInternalServerErrorErr("Não foi possível buscar o usuário", gorm.ErrInvalidTransaction),
				MockValue: gorm.ErrInvalidTransaction,
			},
		},
	}

	s := utest.NewRepositorySuite[UserDomainRepository, model.UserDomainInterface](t, NewUserDomain)

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			var rows = new(sqlmock.Rows)
			if tt.want != nil {
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
						"us_verified",
					}).
					AddRow(
						tt.want.GetID(),
						now,
						now,
						nil,
						tt.want.GetFirstName(),
						tt.want.GetLastName(),
						tt.want.GetEmail(),
						tt.want.GetUsername(),
						tt.want.GetPassword(),
						tt.want.GetVerified(),
					)
			}

			s.Mock.
				ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "us_users" WHERE "us_username" = $1 AND "us_users"."us_deleted_at" IS NULL ORDER BY "us_users"."us_id" LIMIT 1`),
				).
				WithArgs(tt.args.username).
				WillReturnError(tt.expectErr.MockValue).
				WillReturnRows(rows)

			userEntity, err := s.Repo.FindOneByUsername(tt.args.username)
			result := entityconverter.UserEntityToDomain(userEntity)

			if tt.expectErr.Expected {
				assert.Zero(s.T(), userEntity)
				assert.Equal(s.T(), tt.expectErr.Value, err)
			} else {
				assert.NoError(s.T(), err)
				assert.Equal(s.T(), tt.want, result)
			}

			s.TearDownTest()
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
			args: args{email: "igor@gmail.com"},
		},
		{
			name: "Falha informando e-mail não existente",
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewNotFoundErr("Usuário não encontrado", gorm.ErrRecordNotFound),
				MockValue: gorm.ErrRecordNotFound,
			},
			args: args{email: "trentini@gmail.com"},
		},
		{
			name: "Falha por instância do banco de dados inválida",
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewsInternalServerErrorErr("Não foi possível buscar o usuário", gorm.ErrInvalidDB),
				MockValue: gorm.ErrInvalidDB,
			},
		},
		{
			name: "Falha por transação inválida",
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewsInternalServerErrorErr("Não foi possível buscar o usuário", gorm.ErrInvalidTransaction),
				MockValue: gorm.ErrInvalidTransaction,
			},
		},
	}

	s := utest.NewRepositorySuite[UserDomainRepository, model.UserDomainInterface](t, NewUserDomain)

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			var rows = new(sqlmock.Rows)
			if tt.want != nil {
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
						"us_verified",
					}).
					AddRow(
						tt.want.GetID(),
						now,
						now,
						nil,
						tt.want.GetFirstName(),
						tt.want.GetLastName(),
						tt.want.GetEmail(),
						tt.want.GetUsername(),
						tt.want.GetPassword(),
						tt.want.GetVerified(),
					)
			}

			s.Mock.
				ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "us_users" WHERE "us_email" = $1 AND "us_users"."us_deleted_at" IS NULL ORDER BY "us_users"."us_id" LIMIT 1`),
				).
				WithArgs(tt.args.email).
				WillReturnError(tt.expectErr.MockValue).
				WillReturnRows(rows)

			userEntity, err := s.Repo.FindOneByEmail(tt.args.email)
			result := entityconverter.UserEntityToDomain(userEntity)

			if tt.expectErr.Expected {
				assert.Zero(s.T(), userEntity)
				assert.Equal(s.T(), tt.expectErr.Value, err)
			} else {
				assert.NoError(s.T(), err)
				assert.Equal(s.T(), tt.want, result)
			}

			s.TearDownTest()
		})
	}
}
