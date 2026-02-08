package route

import (
	"gorm.io/gorm"
	"kasir-app/handler"
	"kasir-app/repository"
	"kasir-app/service"

	"github.com/gin-gonic/gin"
)

func InitProductRoute(r *gin.Engine, gormDb *gorm.DB) {
	productRepo := repository.NewProductRepository(gormDb)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	r.GET("/products", productHandler.GetAllProduct)
	r.POST("/product", productHandler.CreateProduct)
	r.PUT("/product", productHandler.UpdateProduct)
	r.GET("/product/:id", productHandler.GetProduct)
	r.DELETE("/product/:id", productHandler.DeleteProduct)
}
