package middlewares

import (
	"github.com/gin-gonic/gin"
	"se-backend/config/seerror"
)

func (md *middlewareImpl) Auth(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")

	s, err := getUserByToken(authorization)
	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), err)
		return
	}

	if !s.userIsVerified() {
		err := seerror.NewForbiddenErr("Conta não verificada", nil)
		c.AbortWithStatusJSON(err.GetCode(), err)
		return
	}

	c.Set("user", s.userDomain)

	c.Next()
}
