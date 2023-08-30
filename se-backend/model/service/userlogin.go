package service

import (
	"se-backend/config/seerror"
	"se-backend/model"
)

func (s *userDomainService) Login(userDomain model.UserDomainInterface) (
	model.UserDomainInterface,
	string,
	seerror.SEError,
) {
	userDomain.EncryptPassword()

	user, err := s.FindOneByUsername(userDomain.GetUsername())
	if err != nil {
		return nil, "", err
	}

	if !user.GetVerified() {
		return nil, "", seerror.NewForbiddenErr("Conta não verificada", nil)
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
