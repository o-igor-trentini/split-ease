package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"se-backend/config/seerror"
	"se-backend/config/selog"
	"se-backend/model"
)

type Session struct {
	User model.UserDomainInterface
}

func GetSession(c *gin.Context) Session {
	user, ok := c.Get("user")
	if !ok {
		selog.Error("[SESSION] Não possui um usuário válido", nil)
	}

	return Session{User: user.(model.UserDomainInterface)}
}

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
