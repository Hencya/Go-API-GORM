package main

import (
	Config "orm/config"
	Routes "orm/routes"
)

func main() {
	Config.InitDB()
	echoApp := Routes.New()
	echoApp.Logger.Fatal(echoApp.Start(":8000"))

}

