package routes

import (
	"github.com/gin-gonic/gin"
)

// newHealth instância as rotas de verificação.
func newHealth(r *gin.RouterGroup) {
	base := r.Group("/health")
	{
		base.GET("", func(c *gin.Context) {
			c.JSON(200, "OK")
		})
	}
}
