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

func (co userImpl) Login(c *gin.Context) {
	var userLoginRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		controller.RespondWithError(c, validations.Validate(err))
		return
	}

	domain := model.NewUserLoginDomain(userLoginRequest.Username, userLoginRequest.Password)

	user, token, err := co.service.Login(domain)
	if err != nil {
		controller.RespondWithError(c, err)
		return
	}

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.UserDomainToResponse(user))
}
