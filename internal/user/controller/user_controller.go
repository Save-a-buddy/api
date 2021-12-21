package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"save-a-buddy-api/config"
	"save-a-buddy-api/internal/user/service"
)

type UserController struct {
	cfg         *config.Config
	userService service.UserService
}

func NewUserController(cfg *config.Config, userService service.UserService) UserController {
	return UserController{cfg: cfg, userService: userService}
}

func (uc UserController) GetUsersList() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.userService.FindUsers()

		if err != nil {
			return c.JSON(http.StatusBadRequest, "Error getting data")
		}
		if len(users) == 0 {
			return c.JSON(http.StatusBadRequest, "Empty list")
		}
		return c.JSON(http.StatusOK, users)
	}
}
