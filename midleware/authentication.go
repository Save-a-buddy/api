package midleware

//func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		token := c.Request().Header.Get("Authorization")
//		_, err := auth.ValidateToken(token)
//		if err != nil {
//			return forbidden(c)
//		}
//
//		return next(c)
//	}
//}

//func AuthFirebaseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		firebaseAuth := c.Get("firebaseAuth").(*authFirebase.Client)
//
//		authToken := c.Request().Header.Get("Authorization")
//		idToken := strings.TrimSpace(strings.Replace(authToken, "Bearer", "", 1))
//
//		if idToken == "" {
//			log.Fatal("Is empty token")
//			return next(c)
//		}
//
//		//verify token
//		token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
//		if err != nil {
//			return forbidden(c)
//		}
//
//		c.Set("UUID", token.UID)
//		return next(c)
//	}
//}

//func FireBaseSetupMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		firebaseAuth := config.SetupFirebaseAuth()
//		c.Set("firebaseAuth", firebaseAuth)
//		return next(c)
//	}
//}

//func forbidden(c echo.Context) error {
//	return c.JSON(http.StatusForbidden, getError())
//}

//func getError() map[string]string {
//	return map[string]string{
//		"error": "Don't have auth",
//	}
//}
