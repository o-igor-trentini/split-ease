package routes

import (
	"github.com/gin-gonic/gin"
	userController "se-backend/controller/user"
)

type InitDependencies struct {
	UserCo userController.UserController
}

func Init(rEngine *gin.Engine, dependencies InitDependencies) {
	rEngine.NoRoute(Cors)
	rg := rEngine.Group("/api")

	newHealth(rg)
	newUser(rg, dependencies.UserCo)
}
