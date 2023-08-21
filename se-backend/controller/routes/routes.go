package routes

import (
	"github.com/gin-gonic/gin"
	"se-backend/controller/routes/middlewares"
	userController "se-backend/controller/user"
)

type InitDependencies struct {
	Middleware middlewares.Middleware
	UserCo     userController.UserController
}

func Init(rEngine *gin.Engine, dependencies InitDependencies) {
	rEngine.NoRoute(dependencies.Middleware.Cors)
	rg := rEngine.Group("/api")

	newHealth(rg)
	newUser(rg, dependencies.Middleware, dependencies.UserCo)
}
