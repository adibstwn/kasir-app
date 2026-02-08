package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kasir-app/handler"
	"kasir-app/repository"
	"kasir-app/service"
)

func InitUserRoute(r *gin.Engine, gormDb *gorm.DB) {
	userRepo := repository.NewUserRepository(gormDb)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.POST("/user", userHandler.Create)

}
