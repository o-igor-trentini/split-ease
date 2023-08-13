package service

import (
	"se-backend/config/seerror"
	"se-backend/model"
	"se-backend/model/repository/entity"
)

func (s userDomainService) Create(userDomain model.UserDomainInterface) seerror.SEError {
	userDomain.EncryptPassword()

	usernameUser, err := s.FindOneByUser(userDomain.GetUsername())
	if err != nil && err.GetErr() != seerror.ErrNotFound {
		return err
	}

	if userDomain.GetUsername() == usernameUser.Username {
		return seerror.NewBadRequestErr("Nome de usuário já cadastrado!", err)
	}

	emailUser, err := s.FindOneByEmail(userDomain.GetEmail())
	if err != nil && err.GetErr() != seerror.ErrNotFound {
		return err
	}

	if userDomain.GetEmail() == emailUser.Email {
		return seerror.NewBadRequestErr("E-mail já cadastrado!", err)
	}

	return s.repository.Create(userDomain)
}

func (s userDomainService) FindOneByUser(username string) (entity.User, seerror.SEError) {
	return s.repository.FindOneByUser(username)
}

func (s userDomainService) FindOneByEmail(email string) (entity.User, seerror.SEError) {
	return s.repository.FindOneByEmail(email)
}
