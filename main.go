package main

import (
	"UserManager/config"
	"UserManager/container"
	"UserManager/db"
	"UserManager/db/migration"
	"UserManager/logger"
	"UserManager/server/v1/routes"

	"UserManager/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	conf, env := config.Load()
	logger := logger.NewLogger(env)

	logger.GetZapLogger().Infof("Loaded this configuration : application." + env + ".yml")

	db := db.ConnectDatabase(logger, conf)
	container := container.NewContainer(db, logger, conf, env)
	migration.CreateDatabase(container)
	routes.InitRoutes(e.Group(""), container)

	middleware.InitLoggerMiddleware(e, container)

	if err := e.Start(":8080"); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
	}
}
