package main

import (
	"fmt"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"se-backend/config"
	"se-backend/config/seenv"
	"se-backend/config/selog"
	"se-backend/controller/routes"
)

func main() {
	config.Setup()

	ginMode := map[string]string{
		seenv.ServerModeDebug:   gin.DebugMode,
		seenv.ServerModeRelease: gin.ReleaseMode,
	}
	gin.SetMode(ginMode[seenv.ENV.ServerMode])
	router := gin.Default()
	router.Use(sentrygin.New(sentrygin.Options{}))

	routes.Init(router.Group("/api"))

	if err := router.Run(fmt.Sprintf("%s:%s", seenv.ENV.ServerHost, seenv.ENV.ServerPort)); err != nil {
		selog.Fatal("Não foi possível inicializar o servidor", err, selog.ServerTag)
	}
}
