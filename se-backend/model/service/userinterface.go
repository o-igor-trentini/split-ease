package service

import (
	"se-backend/config/seerror"
	"se-backend/model"
	"se-backend/model/repository"
	"se-backend/model/repository/entity"
)

type UserDomainService interface {
	Create(userDomain model.UserDomainInterface) seerror.SEError
	FindOneByUser(username string) (entity.User, seerror.SEError)
	FindOneByEmail(email string) (entity.User, seerror.SEError)
}

type userDomainService struct {
	repository repository.UserDomainRepository
}

func NewUserDomain(repository repository.UserDomainRepository) UserDomainService {
	return &userDomainService{repository}
}
