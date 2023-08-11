package entityconverter

import (
	"se-backend/model"
	"se-backend/model/repository/entity"
)

func UserDomainToEntity(userDomain model.UserDomainInterface) entity.UserEntity {
	return entity.UserEntity{
		ID:        userDomain.GetID(),
		FirstName: userDomain.GetFirstName(),
		LastName:  userDomain.GetLastName(),
		Email:     userDomain.GetEmail(),
		Username:  userDomain.GetUsername(),
	}
}

func UserEntityToDomain(userEntity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		userEntity.FirstName,
		userEntity.LastName,
		userEntity.Email,
		userEntity.Username,
		userEntity.Password,
	)
	domain.SetID(userEntity.ID)

	return domain
}
