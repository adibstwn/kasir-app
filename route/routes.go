package route

import (
	"database/sql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB, gormDb *gorm.DB) {

	InitProductRoute(r, gormDb)
	InitCategoryRoute(r, db)
	InitUserRoute(r, gormDb)
	InitTransactionRoute(r, gormDb)

}
