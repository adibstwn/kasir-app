package main

import (
	"kasir-app/config"
	"kasir-app/database"
	"kasir-app/route"

	"github.com/gin-gonic/gin"
)

func main() {
	loadConfig, err := config.LoadConfig(".env")
	if err != nil {
		return
	}
	db, err := database.InitDB(loadConfig.DbConnection)
	if err != nil {
		return
	}

	router := gin.Default()
	route.RegisterRoutes(router, db)

	err = router.Run(":" + loadConfig.Port)
	if err != nil {
		return
	}
}
