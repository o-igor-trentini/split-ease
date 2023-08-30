package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"se-backend/config/seerror"
	"se-backend/model"
	"se-backend/model/repository/entity/entityconverter"
	"se-backend/utils/utest"
	"testing"
)

func TestUserDomainRepository_Create(t *testing.T) {
	type args struct {
		userID uint64
		user   model.UserDomainInterface
	}

	type test struct {
		name      string
		want      model.UserDomainInterface
		expectErr utest.Error
		args      args
	}

	newDefaultUser := func() model.UserDomainInterface {
		userDomain := model.NewUserDomain(
			"Igor",
			"Trentini",
			"igor@gmail.com",
			"igor",
			"Senha@123",
		)
		userDomain.EncryptPassword()

		return userDomain
	}

	table := []test{
		{
			name: "Insere com sucesso",
			want: func() model.UserDomainInterface {
				userDomain := newDefaultUser()
				userDomain.SetID(1)

				return userDomain
			}(),
			args: args{userID: 1, user: newDefaultUser()},
		},
		{
			name: "Falha ao tentar inserir com nome de usuário já existente",
			want: newDefaultUser(),
			args: args{userID: 1, user: newDefaultUser()},
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewsInternalServerErrorErr("Não foi possível criar o usuário", gorm.ErrDuplicatedKey),
				MockValue: gorm.ErrDuplicatedKey,
			},
		},
		{
			name: "Falha ao tentar inserir com e-mail já existente",
			want: newDefaultUser(),
			args: args{userID: 1, user: newDefaultUser()},
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewsInternalServerErrorErr("Não foi possível criar o usuário", gorm.ErrDuplicatedKey),
				MockValue: gorm.ErrDuplicatedKey,
			},
		},
		{
			name: "Falha por instância do banco de dados inválida",
			want: newDefaultUser(),
			args: args{userID: 1, user: newDefaultUser()},
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewsInternalServerErrorErr("Não foi possível criar o usuário", gorm.ErrInvalidDB),
				MockValue: gorm.ErrInvalidDB,
			},
		},
		{
			name: "Falha por transação inválida",
			want: newDefaultUser(),
			args: args{userID: 1, user: newDefaultUser()},
			expectErr: utest.Error{
				Expected:  true,
				Value:     seerror.NewsInternalServerErrorErr("Não foi possível criar o usuário", gorm.ErrInvalidTransaction),
				MockValue: gorm.ErrInvalidTransaction,
			},
		},
	}

	s := utest.NewRepositorySuite[UserDomainRepository, model.UserDomainInterface](t, NewUserDomain)

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			s.Mock.ExpectBegin()
			s.Mock.
				ExpectQuery(`INSERT INTO "us_users" (.+) VALUES (.+) RETURNING "us_id"`).
				WillReturnError(tt.expectErr.MockValue).
				WillReturnRows(sqlmock.NewRows([]string{"us_id"}).AddRow(tt.args.userID))

			if tt.expectErr.Expected {
				s.Mock.ExpectRollback()
			} else {
				s.Mock.ExpectCommit()
			}

			userEntity, err := s.Repo.Create(tt.args.user)
			result := entityconverter.UserEntityToDomain(userEntity)

			if tt.expectErr.Expected {
				assert.Equal(s.T(), tt.expectErr.Value, err)
				assert.Equal(s.T(), tt.want, result)
				assert.True(s.T(), 0 == result.GetID())
			} else {
				assert.NoError(s.T(), err)
				assert.Equal(s.T(), tt.want, result)
				assert.NotZero(s.T(), userEntity.CreatedAt)
				assert.NotZero(s.T(), userEntity.UpdatedAt)
				assert.Zero(s.T(), userEntity.DeletedAt)
			}

			s.TearDownTest()
		})
	}
}
