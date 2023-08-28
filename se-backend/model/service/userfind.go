package service

import (
	"se-backend/config/seerror"
	"se-backend/model"
	"se-backend/model/repository/entity/entityconverter"
)

func (s *userDomainService) FindOneByUsername(username string) (model.UserDomainInterface, seerror.SEError) {
	user, err := s.repository.FindOneByUsername(username)
	return entityconverter.UserEntityToDomain(user), err
}

func (s *userDomainService) FindOneByEmail(email string) (model.UserDomainInterface, seerror.SEError) {
	user, err := s.repository.FindOneByEmail(email)
	return entityconverter.UserEntityToDomain(user), err
}
