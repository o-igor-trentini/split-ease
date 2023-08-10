package routes

import "github.com/gin-gonic/gin"

func Init(r *gin.RouterGroup) {
	newHealth(r)
	newUser(r)
}
