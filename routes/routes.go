package routes

import (
	"aysf/day6r1/constants"
	"aysf/day6r1/controllers"
	"aysf/day6r1/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	// e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.POST("/login", controllers.LoginUsersController)

	
	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users/:id", controllers.GetUserDetailControllers)

	eAuth := e.Group("")
	eAuth.Use(middleware.BasicAuth(middlewares.BasicAuthDB))
	eAuth.GET("/users/:id", controllers.GetUserController)
	return e

}
