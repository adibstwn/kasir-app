package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kasir-app/handler"
	"kasir-app/repository"
	"kasir-app/service"
)

// create transaction
//each orderItem insert into transaction detail

func InitTransactionRoute(r *gin.Engine, gormDb *gorm.DB) {
	productRepository := repository.NewProductRepository(gormDb)
	transactionDetailRepository := repository.NewTransactionDetailRepository(gormDb)
	transactionRepository := repository.NewTransactionRepository(gormDb)

	transactionService := service.NewTransactionServiceImpl(gormDb, transactionRepository, transactionDetailRepository, productRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	r.POST("/api/checkout", transactionHandler.Checkout)
	r.GET("/api/report/hari-ini", transactionHandler.ReportToday)
	r.GET("/api/report", transactionHandler.ReportByDate)
}
