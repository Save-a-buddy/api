package server

import (
	"github.com/labstack/echo/v4"
	"save-a-buddy-api/config"
)

type Server struct {
	echo   *echo.Echo
	config *config.Config
}

func New(echo *echo.Echo, config *config.Config) *Server {
	return &Server{echo: echo, config: config}
}

func (s Server) RunServer() error {
	s.echo.Logger.Fatal(s.echo.Start(string(s.config.Server.Port)))
	return nil
}
