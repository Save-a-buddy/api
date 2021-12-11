package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"save-a-buddy-api/config"
	"save-a-buddy-api/internal/auth"
	"save-a-buddy-api/internal/server"
	"save-a-buddy-api/model"
	"save-a-buddy-api/pkg/utils"
)

func main() {
	e := echo.New()
	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		e.Logger.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		e.Logger.Fatalf("ParseConfig: %v", err)
	}

	err = auth.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		e.Logger.Fatalf("Can't be loaded certificates: %v", err)
	}
	s := server.New(e, cfg)
	if err := s.RunServer(); err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Done")
	})

	e.POST("/auth", func(c echo.Context) error {
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
	})


}

func isValidToken(data *model.Login) bool {
	return data.Email == "ae@test.com" && data.Password == "1234"
}
