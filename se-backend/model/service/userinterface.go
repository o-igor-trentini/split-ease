package service

import (
	"se-backend/config/seerror"
	"se-backend/model"
	"se-backend/model/repository"
)

type UserDomainService interface {
	Create(userDomain model.UserDomainInterface) seerror.SEError
}

type userDomainService struct {
	repository repository.UserDomainRepository
}

func NewUserDomain(repository repository.UserDomainRepository) UserDomainService {
	return &userDomainService{repository}
}
