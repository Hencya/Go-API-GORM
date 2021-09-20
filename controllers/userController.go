package Controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	Databases "orm/lib/database"
)

func GetUsersController(ctx echo.Context)error{
	users,err := Databases.GetUsers()
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "success get all users",
		"data": users,
	})
}

func CreateUserController(ctx echo.Context) error{
	user,err := Databases.PostUser(ctx)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "success create new user",
		"data": user,
	})
}