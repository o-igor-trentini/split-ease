package repository

import (
	"se-backend/config/seerror"
	"se-backend/model"
	"se-backend/model/repository/entity/entityconverter"
)

func (r userDomainRepository) Create(userDomain model.UserDomainInterface) seerror.SEError {
	user := entityconverter.UserDomainToEntity(userDomain)

	if err := r.db.Create(&user).Error; err != nil {
		return seerror.NewsInternalServerErrorErr("Não foi possível criar o usuário", err)
	}

	userDomain.SetID(user.ID)

	return nil
}
