package Databases

import (
	"github.com/labstack/echo/v4"
	Config "orm/config"
	Models "orm/models"
)

func GetUsers() (interface{},error){
	var users []Models.Users
	if err := Config.DB.Find(&users).Error;err != nil{
		return nil,err
	}
	return users,nil
}

func PostUser(ctx echo.Context)(interface{},error){
	user := Models.Users{}
	ctx.Bind(&user)

	if err := Config.DB.Save(&user).Error; err != nil{
		return nil,err
	}

	return  user,nil
}
