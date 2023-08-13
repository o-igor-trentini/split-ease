package repository

import (
	"gorm.io/gorm"
	"se-backend/config/seerror"
	"se-backend/model"
	"se-backend/model/repository/entity"
)

type UserDomainRepository interface {
	Create(userDomain model.UserDomainInterface) seerror.SEError
	FindOneByUser(username string) (entity.User, seerror.SEError)
	FindOneByEmail(email string) (entity.User, seerror.SEError)
}

type userDomainRepository struct {
	db *gorm.DB
}

func NewUserDomain(db *gorm.DB) UserDomainRepository {
	return &userDomainRepository{db}
}
