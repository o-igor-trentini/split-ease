package repository

import (
	"se-backend/config/seerror"
	"se-backend/model"
	"se-backend/model/repository/entity"
	"se-backend/model/repository/entity/entityconverter"
)

func (r *userDomainRepository) Create(userDomain model.UserDomainInterface) (entity.User, seerror.SEError) {
	user := entityconverter.UserDomainToEntity(userDomain)

	if err := r.db.Create(&user).Error; err != nil {
		return user, seerror.NewsInternalServerErrorErr("Não foi possível criar o usuário", err)
	}

	return user, nil
}
