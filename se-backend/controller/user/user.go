package controller

import (
	"github.com/gin-gonic/gin"
	"se-backend/model/service"
)

type UserController interface {
	Login(c *gin.Context)
	Create(c *gin.Context)
}

type userImpl struct {
	service               service.UserDomainService
	userActivationService service.UserActivationDomainService
}

// NewUser instância o controller de usuário.
func NewUser(
	service service.UserDomainService,
	userActivationService service.UserActivationDomainService,
) UserController {
	return &userImpl{service, userActivationService}
}
