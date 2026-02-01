package route

import (
	"database/sql"
	"kasir-app/handler"
	"kasir-app/repository"
	"kasir-app/service"

	"github.com/gin-gonic/gin"
)

func InitProductRoute(r *gin.Engine, db *sql.DB) {
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	r.GET("/products", productHandler.GetAllProduct)
	r.POST("/product", productHandler.CreateProduct)
	r.PUT("/product", productHandler.UpdateProduct)
	r.GET("/product/:id", productHandler.GetProduct)
	r.DELETE("/product/:id", productHandler.DeleteProduct)
}
