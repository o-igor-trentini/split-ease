package repository

import (
	"gorm.io/gorm"
	"se-backend/config/seerror"
	"se-backend/model"
)

type UserDomainRepository interface {
	Create(userDomain model.UserDomainInterface) seerror.SEError
}

type userDomainRepository struct {
	db *gorm.DB
}

func NewUserDomain(db *gorm.DB) UserDomainRepository {
	return &userDomainRepository{db}
}
