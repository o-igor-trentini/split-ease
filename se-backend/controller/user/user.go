package controller

import (
	"github.com/gin-gonic/gin"
	"se-backend/model/service"
)

type UserController interface {
	Create(c *gin.Context)
}

type userImpl struct {
	service service.UserDomainService
}

// NewUser instância o controller de usuário.
func NewUser(service service.UserDomainService) UserController {
	return &userImpl{service}
}
