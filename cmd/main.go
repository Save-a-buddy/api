package main

import (
	"github.com/labstack/echo/v4"
	"os"
	"save-a-buddy-api/config"
	"save-a-buddy-api/internal/auth"
	"save-a-buddy-api/internal/server"
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
}
