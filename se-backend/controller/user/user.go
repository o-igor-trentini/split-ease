package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Create(c *gin.Context)
}

type userImpl struct {
}

// NewUser instância o controller de usuário.
func NewUser() UserController {
	return &userImpl{}
}
