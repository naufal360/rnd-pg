package controllers

import (
	"net/http"
	"payment-gateway/dto"
	"payment-gateway/services"
	"payment-gateway/util"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) CreateUser(ctx echo.Context) error {
	var input dto.CreateUserDTO
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.NewErrorResponse("Invalid input"))
	}

	user, err := c.service.CreateUser(input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.NewErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, util.NewSuccessResponse(user))
}

func (c *UserController) GetUsers(ctx echo.Context) error {
	users, err := c.service.GetUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.NewErrorResponse(err.Error()))
	}
	return ctx.JSON(http.StatusOK, util.NewSuccessResponse(users))
}
