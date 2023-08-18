package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"se-backend/config/seenv"
	"se-backend/config/seerror"
	"strings"
	"time"
)

type UserDomainInterface interface {
	GetID() uint64
	GetFirstName() string
	GetLastName() string
	GetEmail() string
	GetUsername() string
	GetPassword() string

	SetID(ID uint64)

	EncryptPassword()
	GenerateToken() (string, seerror.SEError)
}

type userDomain struct {
	id        uint64
	firstName string
	lastName  string
	email     string
	username  string
	password  string
}

func NewUserDomain(firstName, lastName, email, username, password string) UserDomainInterface {
	caser := cases.Title(language.Und)

	return &userDomain{
		firstName: caser.String(strings.TrimSpace(firstName)),
		lastName:  caser.String(strings.TrimSpace(lastName)),
		email:     strings.ToLower(strings.TrimSpace(email)),
		username:  strings.TrimSpace(username),
		password:  password,
	}
}

func NewUserLoginDomain(username, password string) UserDomainInterface {
	return &userDomain{
		username: strings.TrimSpace(username),
		password: password,
	}
}

func (ud *userDomain) GetID() uint64 {
	return ud.id
}

func (ud *userDomain) GetFirstName() string {
	return ud.firstName
}

func (ud *userDomain) GetLastName() string {
	return ud.lastName
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetUsername() string {
	return ud.username
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) SetID(ID uint64) {
	ud.id = ID
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))

	ud.password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) GenerateToken() (string, seerror.SEError) {
	claims := jwt.MapClaims{
		"id":        ud.id,
		"firstName": ud.firstName,
		"lastName":  ud.lastName,
		"email":     ud.email,
		"username":  ud.username,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte(seenv.ENV.JWTSecretKey))
	if err != nil {
		return strToken, seerror.NewsInternalServerErrorErr("Não foi possível gerar o token", err)
	}

	return strToken, nil
}

func VerifyToken(value string) (UserDomainInterface, seerror.SEError) {
	removeBearerPrefix := func(token string) string {
		prefix := "Bearer "

		if !strings.HasPrefix(token, prefix) {
			return token
		}

		return strings.TrimPrefix(token, prefix)
	}

	token, err := jwt.Parse(removeBearerPrefix(value), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(seenv.ENV.JWTSecretKey), nil
		}

		return nil, seerror.NewBadRequestErr("Token inválido", nil)
	})
	if err != nil {
		return nil, seerror.NewUnauthorizedErr("Token inválido", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, seerror.NewUnauthorizedErr("Token inválido", err)
	}

	return &userDomain{
		id:        uint64(claims["id"].(float64)),
		firstName: claims["firstName"].(string),
		lastName:  claims["lastName"].(string),
		email:     claims["email"].(string),
		username:  claims["username"].(string),
	}, nil
}
