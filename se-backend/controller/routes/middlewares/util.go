package middlewares

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"se-backend/config/seenv"
	"se-backend/config/seerror"
	"se-backend/model"
	"strings"
)

func getUserByToken(authorization string) (session, seerror.SEError) {
	var s session

	removeBearerPrefix := func(token string) string {
		prefix := "Bearer "

		if !strings.HasPrefix(token, prefix) {
			return token
		}

		return strings.TrimPrefix(token, prefix)
	}

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

		return s, seerror.NewUnauthorizedErr(message, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return s, seerror.NewUnauthorizedErr("Token inválido", err)
	}

	userDomain := model.NewUserDomain(
		claims["firstName"].(string),
		claims["lastName"].(string),
		claims["email"].(string),
		claims["username"].(string),
		"",
	)
	userDomain.SetID(uint64(claims["id"].(float64)))
	userDomain.SetVerified(claims["verified"].(bool))

	s.userDomain = userDomain
	s.token = authorization

	return s, nil
}
