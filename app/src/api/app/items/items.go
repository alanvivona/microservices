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
	group := r.Group("/item")
	{
		group.GET("", GetItems)
		group.GET("/:id", GetItem)
		group.POST("/", PostItem)
		group.DELETE("/:id", DeleteItem)
	}
}
