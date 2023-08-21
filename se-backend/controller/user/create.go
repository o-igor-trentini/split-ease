package controller

import (
	"github.com/gin-gonic/gin"
	"se-backend/config/validations"
	"se-backend/controller"
	"se-backend/controller/model/request"
	"se-backend/model"
	"se-backend/view"
)

func (co userImpl) Create(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		controller.RespondWithError(c, validations.Validate(err))
		return
	}

	domain := model.NewUserDomain(
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.Email,
		userRequest.Username,
		userRequest.Password,
	)

	if err := co.service.Create(domain); err != nil {
		controller.RespondWithError(c, err)
		return
	}

	// TODO: Criar verificação de conta do usuário por email.

	c.JSON(201, view.UserDomainToResponse(domain))
}
