package Routes

import (
	"github.com/labstack/echo/v4"
	Controller "orm/controllers"
)

func New() *echo.Echo{
	echoApp := echo.New()

	echoApp.GET("/users",Controller.GetUsersController)
	echoApp.POST("/user",Controller.CreateUserController)

	return echoApp
}
