package main

import (
	"kasir-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)

	err := router.Run(":8081")
	if err != nil {
		return
	} // listens on 0.0.0.0:8080 by default
}
