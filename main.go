package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var (
	DB *gorm.DB
)

//auto migrate doesnt have rollback
type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}


func main() {
	Init()

	echoApp := echo.New()

	echoApp.GET("/users",getUsersController)
	echoApp.POST("/user",createUserController)

	echoApp.Start(":8000")

}

func Init(){
	InitDB()
	initialMigration()
}

func InitDB(){
	var err error
	connectingString := "root:root@/coba_echo?charset=utf8&parseTime=True&loc=Local"
	DB,err = gorm.Open(mysql.Open(connectingString),&gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func getUsersController(ctx echo.Context)error{
	var users []User
	if err := DB.Find(&users).Error;err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "success get all users",
		"data": users,
	})
}

func createUserController(ctx echo.Context) error{
	user := User{}
	ctx.Bind(&user)

	if err := DB.Save(&user).Error; err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "success create new user",
		"data": user,
	})
}

func initialMigration(){
	DB.AutoMigrate(&User{})
}
