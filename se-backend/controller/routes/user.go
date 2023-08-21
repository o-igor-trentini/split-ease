package routes

import (
	"github.com/gin-gonic/gin"
	"se-backend/controller/routes/middlewares"
	controller "se-backend/controller/user"
)

// newUser instância as rotas de usuário.
func newUser(r *gin.RouterGroup, md middlewares.Middleware, co controller.UserController) {
	base := r.Group("/user", md.Cors)
	{
		base.POST("/login", co.Login)
		base.POST("", md.Auth, co.Create)
	}
}
