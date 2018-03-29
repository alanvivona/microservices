package gdrive

import (
	"api/app/models"

	"github.com/gin-gonic/gin"
)

//GdS google drive service
var (
	Gds models.GdriveServiceInterface
)

// Configure for google drive
func Configure(r *gin.Engine) {
	Gds = &GdriveService{}
	group := r.Group("/gdrive")
	{
		group.GET("/auth", Auth)
		group.GET("/search-in-doc/:id", SearchInDoc)
		group.POST("/file", CreateFile)
	}
}
