package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// newHealth instância as rotas de verificação.
func newHealth(r *gin.RouterGroup) {
	base := r.Group("/health")
	{
		base.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, "OK")
		})
	}
}
