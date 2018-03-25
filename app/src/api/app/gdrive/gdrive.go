package gdrive

import (
	"api/app/models"

	"github.com/gin-gonic/gin"
)

//GdS google drive service
var (
	GdService models.GdriveServiceInterface
)

// Configure for google drive
func Configure(r *gin.Engine) {
	Gds = &GdriveService{}

	r.GET("/search-in-doc/:id", SearchInDoc)
	r.POST("/file", CreateFile)
}
