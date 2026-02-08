package main

import (
	"kasir-app/config"
	"kasir-app/database"
	"kasir-app/middleware"
	"kasir-app/route"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig(".env")

	//database.InitDB(config.AppConfig.DbConnection)
	database.InitDBGorm()

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.ResponseMiddleware())

	route.RegisterRoutes(router, database.SqlDB, database.GormDB)

	err := router.Run(":" + config.AppConfig.Port)
	if err != nil {
		return
	}
}
