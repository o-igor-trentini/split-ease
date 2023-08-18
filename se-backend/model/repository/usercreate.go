package repository

import (
	"se-backend/config/seerror"
	"se-backend/model"
	"se-backend/model/repository/entity"
	"se-backend/model/repository/entity/entityconverter"
	"se-backend/utils/ugorm"
)

func (r userDomainRepository) Create(userDomain model.UserDomainInterface) (entity.User, seerror.SEError) {
	user := entityconverter.UserDomainToEntity(userDomain)

	if err := r.db.Create(&user).Error; err != nil {
		return user, seerror.NewsInternalServerErrorErr("Não foi possível criar o usuário", err)
	}

	return user, nil
}

func (r userDomainRepository) FindOneByUsername(username string) (entity.User, seerror.SEError) {
	var user entity.User

	if err := r.db.Where("us_username", username).First(&user).Error; err != nil {
		if ugorm.IsErrNotFound(err) {
			return user, seerror.NewNotFoundErr("Usuário não encontrado", err)
		}

		return user, seerror.NewsInternalServerErrorErr("Não foi possível buscar o usuário", err)
	}

	return user, nil
}

func (r userDomainRepository) FindOneByEmail(email string) (entity.User, seerror.SEError) {
	var user entity.User

	if err := r.db.Where("us_email", email).First(&user).Error; err != nil {
		if ugorm.IsErrNotFound(err) {
			return user, seerror.NewNotFoundErr("Usuário não encontrado", err)
		}

		return user, seerror.NewsInternalServerErrorErr("Não foi possível buscar o usuário", err)
	}

	return user, nil
}
