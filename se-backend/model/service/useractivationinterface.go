package service

import (
	"se-backend/config/seerror"
	"se-backend/model"
	"se-backend/model/repository"
)

type UserActivationDomainService interface {
	Create(userDomain model.UserDomainInterface) seerror.SEError
}

type userActivationDomainService struct {
	repository repository.UserActivationDomainInterface
}

func NewUserActivationDomain(repository repository.UserActivationDomainInterface) UserActivationDomainService {
	return &userActivationDomainService{repository}
}
