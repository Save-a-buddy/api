package server

import (
	"github.com/labstack/echo/v4"
	"save-a-buddy-api/internal/enrollment/controller"
)

func (s *Server) HandlerRoute(e *echo.Echo) error {
	apiGroup := s.echo.Group("/api/v1")
	authGroup := apiGroup.Group("/user")

	//LoginController
	lc := controller.LoginControllerNew(s.config)
	authGroup.POST("/login", lc.Login())

	return nil
}
