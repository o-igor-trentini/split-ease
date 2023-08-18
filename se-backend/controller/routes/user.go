package routes

import (
	"github.com/gin-gonic/gin"
	controller "se-backend/controller/user"
)

// newUser instância as rotas de usuário.
func newUser(r *gin.RouterGroup, co controller.UserController) {
	base := r.Group("/user", Cors)
	{
		base.POST("/login", co.Login)
		base.POST("", co.Create)
	}
}
