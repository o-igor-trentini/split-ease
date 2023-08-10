package sesentry

import "encoding/json"

type LoggerPayload struct {
	Time       string  `json:"time"`
	Level      string  `json:"level"`
	Tag        *string `json:"tag"`
	Message    string  `json:"message"`
	Function   string  `json:"function"`
	StackTrace string  `json:"stackTrace"`
}

func (l *LoggerPayload) JSONToStruct(p []byte) {
	_ = json.Unmarshal(p, l)
}
