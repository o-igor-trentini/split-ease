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
	// TODO: Implementar refresh token

	claims := jwt.MapClaims{
		"id":        ud.id,
		"firstName": ud.firstName,
		"lastName":  ud.lastName,
		"email":     ud.email,
		"username":  ud.username,
		"exp":       time.Now().Add(5 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte(seenv.ENV.JWTSecretKey))
	if err != nil {
		return strToken, seerror.NewsInternalServerErrorErr("Não foi possível gerar o token", err)
	}

	return strToken, nil
}
