package seerror

import "se-backend/config/selog"

type SEError interface {
	Error() string
	GetCode() int
	GetMessage() string
	GetErr() string
	GetCauses() []Cause
	GetSource() error
}

type seErrorImpl struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Causes  []Cause `json:"causes"`
	Source  error   `json:"-"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (s seErrorImpl) Error() string {
	return s.Message
}

func (s seErrorImpl) GetCode() int {
	return s.Code
}

func (s seErrorImpl) GetMessage() string {
	return s.Message
}

func (s seErrorImpl) GetErr() string {
	return s.Err
}

func (s seErrorImpl) GetCauses() []Cause {
	return s.Causes
}

func (s seErrorImpl) GetSource() error {
	return s.Source
}

func New(code int, message, err string, causes []Cause, source error) SEError {
	v := &seErrorImpl{
		Code:    code,
		Message: message,
		Err:     err,
		Causes:  causes,
		Source:  source,
	}

	go selog.Error(v.GetMessage(), v)

	return v
}
