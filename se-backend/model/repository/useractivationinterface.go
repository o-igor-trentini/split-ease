package repository

import (
	"gorm.io/gorm"
)

type UserActivationDomainInterface interface {
}

type userActivationDomainRepository struct {
	db *gorm.DB
}

func NewUserActivationDomain(db *gorm.DB) UserActivationDomainInterface {
	return &userActivationDomainRepository{db}
}
