package route

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {

	InitProductRoute(r, db)
	InitCategoryRoute(r, db)

}
