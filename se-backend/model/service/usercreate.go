package service

import (
	"se-backend/config/seerror"
	"se-backend/model"
)

func (s userDomainService) Create(userDomain model.UserDomainInterface) seerror.SEError {
	userDomain.EncryptPassword()

	return s.repository.Create(userDomain)
}
