package controller

import (
	"github.com/gin-gonic/gin"
	"se-backend/config/validations"
	"se-backend/controller"
	request "se-backend/controller/model/request/user"
)

func (co userImpl) Create(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		controller.RespondWithError(c, validations.Validate(err))
		return
	}

	c.JSON(200, "Criado")
}
