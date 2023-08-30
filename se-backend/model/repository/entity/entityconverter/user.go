package entityconverter

import (
	"se-backend/model"
	"se-backend/model/repository/entity"
)

func UserDomainToEntity(userDomain model.UserDomainInterface) entity.User {
	return entity.User{
		ID:        userDomain.GetID(),
		FirstName: userDomain.GetFirstName(),
		LastName:  userDomain.GetLastName(),
		Email:     userDomain.GetEmail(),
		Username:  userDomain.GetUsername(),
		Password:  userDomain.GetPassword(),
		Verified:  userDomain.GetVerified(),
	}
}

func UserEntityToDomain(userEntity entity.User) model.UserDomainInterface {
	domain := model.NewUserDomain(
		userEntity.FirstName,
		userEntity.LastName,
		userEntity.Email,
		userEntity.Username,
		userEntity.Password,
	)
	domain.SetID(userEntity.ID)
	domain.SetVerified(userEntity.Verified)

	return domain
}
