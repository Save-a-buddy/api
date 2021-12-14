package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"save-a-buddy-api/config"
	"save-a-buddy-api/internal/auth"
	"save-a-buddy-api/model"
)

type LoginController struct {
	cfg *config.Config
}

func NewLoginController(cfg *config.Config) LoginController {
	return LoginController{cfg: cfg}
}

func (lc LoginController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := new(model.Login)
		if err := c.Bind(data); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if !isValidToken(data) {
			return echo.NewHTTPError(http.StatusBadRequest, "User or Password invalid")
		}

		token, err := auth.GenerateToken(data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		response := make(map[string]string)
		response["token"] = token

		return c.JSON(http.StatusOK, response)

	}
}

func isValidToken(data *model.Login) bool {
	return data.Email == "ae@test.com" && data.Password == "1234"
}
