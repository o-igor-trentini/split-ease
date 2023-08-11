package main

import (
	"fmt"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"se-backend/config"
	"se-backend/config/database/sepostgres"
	"se-backend/config/database/sepostgres/sepostgresmigration"
	"se-backend/config/seenv"
	"se-backend/config/selog"
	"se-backend/controller/routes"
	controller "se-backend/controller/user"
	"se-backend/model/repository"
	"se-backend/model/service"
)

func main() {
	config.Setup()
	db := sepostgres.New()
	args(db)

	ginMode := map[string]string{
		seenv.ServerModeDebug:   gin.DebugMode,
		seenv.ServerModeRelease: gin.ReleaseMode,
	}
	gin.SetMode(ginMode[seenv.ENV.ServerMode])
	router := gin.Default()
	router.Use(sentrygin.New(sentrygin.Options{}))
	routes.Init(router, initDependencies(db))

	if err := router.Run(fmt.Sprintf("%s:%s", seenv.ENV.ServerHost, seenv.ENV.ServerPort)); err != nil {
		selog.Fatal("Não foi possível inicializar o servidor", err, selog.ServerTag)
	}
}

func initDependencies(db *gorm.DB) routes.InitDependencies {
	userDomainRepository := repository.NewUserDomain(db)
	userDomainService := service.NewUserDomain(userDomainRepository)

	return routes.InitDependencies{
		UserCo: controller.NewUser(userDomainService),
	}
}

func args(db *gorm.DB) {
	if len(os.Args) <= 1 {
		return
	}

	for _, arg := range os.Args {
		switch arg {
		case "-cds", "--create-database-structure":
			sepostgresmigration.CreateDatabaseStructure(db)
			sepostgresmigration.ExecMigrations(db)

		case "-cm", "--create-migration":
			sepostgresmigration.CreateMigrationFile()

		case "-m", "--migrate":
			sepostgresmigration.ExecMigrations(db)

		case "-mr", "--migrate-and-run":
			sepostgresmigration.ExecMigrations(db)
			return
		}
	}

	os.Exit(0)
}
