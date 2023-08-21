package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"se-backend/config/seenv"
	"se-backend/config/seerror"
	"se-backend/model"
	"strings"
)

func (md *middlewareImpl) Auth(c *gin.Context) {
	removeBearerPrefix := func(token string) string {
		prefix := "Bearer "

		if !strings.HasPrefix(token, prefix) {
			return token
		}

		return strings.TrimPrefix(token, prefix)
	}

	authorization := c.Request.Header.Get("Authorization")

	token, err := jwt.Parse(removeBearerPrefix(authorization), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(seenv.ENV.JWTSecretKey), nil
		}

		return nil, seerror.NewBadRequestErr("Token inválido", nil)
	})
	if err != nil {
		message := "Token inválido"
		if errors.Is(err, jwt.ErrTokenExpired) {
			message = "Token expirado"
		}

		err := seerror.NewUnauthorizedErr(message, err)
		c.AbortWithStatusJSON(err.GetCode(), err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err := seerror.NewUnauthorizedErr("Token inválido", err)
		c.AbortWithStatusJSON(err.GetCode(), err)
		return
	}

	user := model.NewUserDomain(
		claims["firstName"].(string),
		claims["lastName"].(string),
		claims["email"].(string),
		claims["username"].(string),
		"",
	)
	user.SetID(uint64(claims["id"].(float64)))

	c.Set("user", user)

	c.Next()
}
