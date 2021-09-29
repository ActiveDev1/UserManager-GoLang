package routes

import (
	"github.com/labstack/echo/v4"

	"UserManager/container"
	"UserManager/server/v1/controller"
)

func InitRoutes(g *echo.Group, container container.Container) {
	UserInit(g.Group(controller.APIUser), container)
}
