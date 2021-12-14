package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"save-a-buddy-api/config"
)

type UserController struct {
	cfg *config.Config
}

func NewUserController(cfg *config.Config) UserController {
	return UserController{cfg: cfg}
}

func (uc UserController) GetUsersList() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Users List")
	}
}
