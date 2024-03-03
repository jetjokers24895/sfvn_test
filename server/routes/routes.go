package routes

import (
	s "sfvn_test/server"
	"sfvn_test/server/handlers"
	customMiddleware "sfvn_test/server/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *s.Server) {
	anonymousHandler := handlers.NewAnonymHandler(server)
	userHandler := handlers.NewUserHandler(server)

	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	server.Echo.GET("/health-check", anonymousHandler.HealthCheck)
	//===user zone===
	server.Echo.GET("/get_histories", userHandler.GetHistories, customMiddleware.IsAuthentication) // have not yet created profile

}
