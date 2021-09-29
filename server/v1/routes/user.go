package routes

import (
	"UserManager/container"
	"UserManager/middleware"
	"UserManager/server/v1/controller"

	"github.com/labstack/echo/v4"
)

func UserInit(g *echo.Group, container container.Container) {
	user := controller.NewUserController(container)
	g.Use(middleware.Authorize)

	g.POST("", user.Create)
	g.GET("", user.GetAll)
	g.PATCH(controller.APIID, user.UpdateOne)
	g.DELETE(controller.APIID, user.DeleteOne)
}
