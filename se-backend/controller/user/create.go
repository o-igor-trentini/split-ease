package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	userDomain := model.NewUserDomain(
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.Email,
		userRequest.Username,
		userRequest.Password,
	)

	if err := co.service.Create(userDomain); err != nil {
		controller.RespondWithError(c, err)
		return
	}

	if err := co.userActivationService.Create(userDomain); err != nil {
		controller.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, view.UserDomainToResponse(userDomain))
}
