package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"se-backend/config/seerror"
	"se-backend/config/selog"
)

// RespondWithError realiza a resposta HTTP de forma semi automática.
func RespondWithError(c *gin.Context, res error) {
	var cErr seerror.SEError
	if errors.As(res, &cErr) {
		c.AbortWithStatusJSON(cErr.GetCode(), cErr)
		return
	}

	selog.Error("Erro desconhecido", res)
	c.AbortWithStatusJSON(500, res)
}
