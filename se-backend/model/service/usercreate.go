package service

import (
	"se-backend/config/seerror"
	"se-backend/model"
)

func (s userDomainService) Create(userDomain model.UserDomainInterface) seerror.SEError {
	userDomain.EncryptPassword()

	usernameUser, err := s.FindOneByUsername(userDomain.GetUsername())
	if err != nil && err.GetErr() != seerror.ErrNotFound {
		return err
	}

	if userDomain.GetUsername() == usernameUser.GetUsername() {
		return seerror.NewBadRequestErr("Nome de usuário já cadastrado!", err)
	}

	emailUser, err := s.FindOneByEmail(userDomain.GetEmail())
	if err != nil && err.GetErr() != seerror.ErrNotFound {
		return err
	}

	if userDomain.GetEmail() == emailUser.GetEmail() {
		return seerror.NewBadRequestErr("E-mail já cadastrado!", err)
	}

	user, err := s.repository.Create(userDomain)
	if err != nil {
		return err
	}

	userDomain.SetID(user.ID)

	return nil
}
