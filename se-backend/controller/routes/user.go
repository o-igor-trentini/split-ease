package routes

import (
	"github.com/gin-gonic/gin"
	controller "se-backend/controller/user"
)

// newUser instância as rotas de usuário.
func newUser(r *gin.RouterGroup) {
	co := controller.NewUser()

	base := r.Group("/user")
	{
		base.POST("", co.Create)
	}
}
