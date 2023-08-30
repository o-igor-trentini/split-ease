package service

import (
	"bytes"
	"fmt"
	"html/template"
	"se-backend/config/seenv"
	"se-backend/config/seerror"
	"se-backend/model"
	"se-backend/third_party/seemail"
	"strings"
	"time"
)

func (s *userActivationDomainService) Create(userDomain model.UserDomainInterface) seerror.SEError {
	// TODO: Criar verificação de conta do usuário por email.
	// - Criar registro de verificação de usuário
	// - Enviar e-mail com código de verificação
	// - Validar e ativar usuário

	go func() { _ = s.sendEmailForActivation(userDomain) }()

	return nil
}

func (s *userActivationDomainService) sendEmailForActivation(userDomain model.UserDomainInterface) seerror.SEError {
	eService := seemail.New(seenv.ENV.EmailAPIKey, "Split Ease", seenv.ENV.EmailFrom)

	url := "./config/emailtemplate/verifyuser.html"
	filePath := strings.Split(url, "/")
	// o nome do arquivo é o último item do caminho
	fileName := filePath[len(filePath)-1]

	t, err := template.New(fileName).ParseFiles(url)
	if err != nil {
		return seerror.NewsInternalServerErrorErr("Não foi possível buscar a estrutura do e-mail", err)
	}

	type templateEmail struct {
		PrimaryColor       string
		Title              string
		Name               string
		Link               string
		ExpirationDatetime string
	}

	// TODO: Alterar para token dinâmico
	// - Utilizar encoder do valor na url
	token := "token_ativacao_123i091238902"

	data := templateEmail{
		PrimaryColor:       "#F2CD60",
		Title:              "Ativação da conta",
		Name:               userDomain.GetFirstName(),
		Link:               fmt.Sprintf("%s/usuario/ativar/%s", seenv.ENV.FrontendURL, token),
		ExpirationDatetime: time.Now().Format("15:04 02/01/2006"),
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return seerror.NewsInternalServerErrorErr("Não foi possível buscar a estrutura do e-mail", err)
	}

	htmlBody := buf.String()

	if err := eService.Send(
		"Verificação da conta",
		userDomain.GetUsername(),
		userDomain.GetEmail(),
		"Vamos ativar a sua conta?",
		htmlBody,
	); err != nil {
		return err
	}

	return nil
}
