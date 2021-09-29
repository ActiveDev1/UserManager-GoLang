package controller

import (
	"UserManager/container"
	"UserManager/model/dto"
	"UserManager/server/v1/service"
	"net/http"

	"UserManager/util"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	container container.Container
	service   *service.UserService
}

func NewUserController(container container.Container) *UserController {
	return &UserController{container: container, service: service.NewUserService(container)}
}

func (controller *UserController) Create(c echo.Context) error {
	dto := dto.NewUserDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}
	user, err := controller.service.CreateNewUser(dto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, user)
}

func (controller *UserController) GetAll(c echo.Context) error {
	users, err := controller.service.FindAllUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func (controller *UserController) UpdateOne(c echo.Context) error {
	dto := dto.NewUserDto()

	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}

	user, err := controller.service.UpdateUser(dto, util.ConvertToUint(c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) DeleteOne(c echo.Context) error {
	err := controller.service.DeleteUser(util.ConvertToUint(c.Param("id")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.NoContent((http.StatusNoContent))
}
