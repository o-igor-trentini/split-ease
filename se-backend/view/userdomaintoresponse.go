package view

import (
	"se-backend/controller/model/response"
	"se-backend/model"
)

func UserDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:        userDomain.GetID(),
		FirstName: userDomain.GetFirstName(),
		LastName:  userDomain.GetLastName(),
		Email:     userDomain.GetEmail(),
		Username:  userDomain.GetUsername(),
	}
}
