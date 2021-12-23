package server

import (
	loginController "save-a-buddy-api/internal/enrollment/controller"
	userController "save-a-buddy-api/internal/user/controller"
	"save-a-buddy-api/internal/user/repository"
	"save-a-buddy-api/internal/user/service"
)

func (s *Server) HandlerRoute() error {
	apiGroup := s.echo.Group("/api/v1")
	userGroup := apiGroup.Group("/user")

	//Controllers

	lc := loginController.NewLoginController(s.config)
	ur := repository.New(s.MongoDb)
	us := service.New(ur)
	uc := userController.NewUserController(s.config, us)

	//Login
	apiGroup.POST("/login", lc.Login())

	//User
	//userGroup.Use(midleware.AuthenticationMiddleware)
	userGroup.GET("", uc.GetUsersList())

	return nil
}
