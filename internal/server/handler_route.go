package server

import (
	loginController "save-a-buddy-api/internal/enrollment/controller"
	userController "save-a-buddy-api/internal/user/controller"
	"save-a-buddy-api/midleware"
)

func (s *Server) HandlerRoute() error {
	apiGroup := s.echo.Group("/api/v1")
	userGroup := apiGroup.Group("/user")

	//Controllers

	lc := loginController.NewLoginController(s.config)
	uc := userController.NewUserController(s.config)

	//Login
	apiGroup.POST("/login", lc.Login())

	//User
	userGroup.Use(midleware.AuthenticationMiddleware)
	userGroup.GET("", uc.GetUsersList())

	return nil
}
