package seerror

type SEError interface {
	Error() string
	GetCode() int
	GetMessage() string
	GetErr() string
	GetCauses() []Cause
}

type seErrorImpl struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Causes  []Cause `json:"causes"`
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
