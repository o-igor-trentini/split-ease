package middlewares

import (
	"se-backend/model"
)

type session struct {
	token      string
	userDomain model.UserDomainInterface
}

func (s *session) userIsVerified() bool {
	return s.userDomain.GetVerified()
}
