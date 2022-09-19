package main

import (
	_ "go-web-template/docs"
	"go-web-template/modules/engine"
	"go-web-template/modules/repository"
)

// @title           Go Web Template
// @version         1.0
// @description     This is a template web API project of Go
// @contact.name   Ray Lin
// @contact.email  ray.lin@shoalter.com
// @host      localhost:8081
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
    ginEngine := engine.InitGinManager().GetGinEngine()

    ginEngine.Run(":8081")

    defer repository.CloseMySQL()
}
