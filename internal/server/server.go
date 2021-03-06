package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"save-a-buddy-api/config"
	"save-a-buddy-api/db"
)

type Server struct {
	echo    *echo.Echo
	config  *config.Config
	MongoDb *db.MongoDb
}

func New(echo *echo.Echo, config *config.Config, mongoDb *db.MongoDb) *Server {
	return &Server{echo: echo, config: config, MongoDb: mongoDb}
}

func (s Server) RunServer() error {
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: s.echo.Logger.Output(),
	}))

	server := &http.Server{
		Addr: s.config.Server.Port,
	}

	if err := s.HandlerRoute(); err != nil {
		return err
	}

	if err := s.echo.StartServer(server); err != nil {
		s.echo.Logger.Printf("Error starting server %s", err)
	}
	return nil
}
