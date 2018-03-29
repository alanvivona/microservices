package gdrive

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth ...
func Auth(c *gin.Context) {

	// Verify if is a redirect from Gdrive with the authorized token
	stateToken := c.Query("state")
	if stateToken != "" {
		// Getting new token from Gdrive
		tokenCode := c.Query("code")
		if stateToken != "state-token" || tokenCode == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "auth_error", "description": "Invalid token"})
		} else {
			// TODO SAVE TOKEN TO SERVICE
			c.JSON(http.StatusOK, gin.H{"success": "auth_success", "description": "Authentication success"})
		}
	} else {
		// First time auth. Provide auth URL to the user
		err := Gds.CreateClient()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "auth_error", "description": "Auth process error"})
		}
		c.JSON(http.StatusOK, gin.H{"auth": "OK", "go to the following URL to authorize the ML Challenge API": "http://LALALALALALALALLA"})
	}
}

// SearchInDoc ...
func SearchInDoc(c *gin.Context) {
	// TODO
	c.JSON(http.StatusInternalServerError, gin.H{"error": "auth_error", "description": "not implemented yet"})
}

// CreateFile ...
func CreateFile(c *gin.Context) {
	// TODO
	c.JSON(http.StatusInternalServerError, gin.H{"error": "auth_error", "description": "not implemented yet"})
}
