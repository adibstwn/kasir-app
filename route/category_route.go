package route

import (
	"database/sql"
	"kasir-app/handler"
	"kasir-app/repository"
	"kasir-app/service"

	"github.com/gin-gonic/gin"
)

func InitCategoryRoute(r *gin.Engine, db *sql.DB) {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	r.GET("/categories", categoryHandler.GetAllCategory)
	r.POST("/categories", categoryHandler.CreateCategory)
	r.PUT("/category", categoryHandler.UpdateCategory)
	r.GET("/category/:id", categoryHandler.GetCategoryById)
	r.DELETE("/category/:id", categoryHandler.DeleteCategory)

}
