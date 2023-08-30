package seemail

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"se-backend/config/seerror"
)

type Service interface {
	Send(subject, emailToName, emailTo, plainTextContent, htmlContent string) seerror.SEError
}

type serviceImpl struct {
	APIKey        string
	EmailFromName string
	EmailFrom     string
}

func New(APIKey, EmailFromName, EmailFrom string) Service {
	return &serviceImpl{
		APIKey:        APIKey,
		EmailFromName: EmailFromName,
		EmailFrom:     EmailFrom,
	}
}

func (s *serviceImpl) Send(subject, emailToName, emailTo, plainTextContent, htmlContent string) seerror.SEError {
	from := mail.NewEmail(s.EmailFromName, s.EmailFrom)
	to := mail.NewEmail(emailToName, emailTo)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(s.APIKey)

	if _, err := client.Send(message); err != nil {
		return seerror.NewBadRequestErr("Não foi possível enviar o e-mail", err)
	}

	return nil
}
