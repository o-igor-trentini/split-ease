package config

import (
	"se-backend/config/seenv"
	"se-backend/config/selog"
)

// Setup executa as configurações do projeto.
func Setup() {
	seenv.Init()
	selog.Init()
}
