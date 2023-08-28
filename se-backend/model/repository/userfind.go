package repository

import (
	"se-backend/config/seerror"
	"se-backend/model/repository/entity"
	"se-backend/utils/ugorm"
)

func (r *userDomainRepository) FindOneByUsername(username string) (entity.User, seerror.SEError) {
	var user entity.User

	if err := r.db.Where("us_username", username).First(&user).Error; err != nil {
		if ugorm.IsErrNotFound(err) {
			return user, seerror.NewNotFoundErr("Usuário não encontrado", err)
		}

		return user, seerror.NewsInternalServerErrorErr("Não foi possível buscar o usuário", err)
	}

	return user, nil
}

func (r *userDomainRepository) FindOneByEmail(email string) (entity.User, seerror.SEError) {
	var user entity.User

	if err := r.db.Where("us_email", email).First(&user).Error; err != nil {
		if ugorm.IsErrNotFound(err) {
			return user, seerror.NewNotFoundErr("Usuário não encontrado", err)
		}

		return user, seerror.NewsInternalServerErrorErr("Não foi possível buscar o usuário", err)
	}

	return user, nil
}
