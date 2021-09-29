package main

import (
	"UserManager/config"
	"UserManager/container"
	"UserManager/db"
	"UserManager/db/migration"
	"UserManager/server/v1/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	conf, env := config.Load()
	db := db.ConnectDatabase(conf)
	container := container.NewContainer(db, conf, env)
	migration.CreateDatabase(container)
	routes.InitRoutes(e.Group(""), container)

	e.Logger.Fatal(e.Start(":8080"))
}
