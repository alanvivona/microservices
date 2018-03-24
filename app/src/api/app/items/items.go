package items

import (
	"api/app/models"
	"database/sql"

	"github.com/gin-gonic/gin"
)

var (
	//Is item service
	Is models.ItemServiceInterface
)

// Configure for items
func Configure(r *gin.Engine, db *sql.DB) {
	Is = &ItemService{DB: db}

	r.GET("/item/:id", GetItem)
	r.POST("/item", PostItem)
	r.DELETE("/item/:id", nil)
	r.GET("/item", nil)
}
