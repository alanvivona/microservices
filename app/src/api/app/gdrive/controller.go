package gdrive

import (
	"api/app/models"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Auth ...
func Auth(c *gin.Context) {

	// Verify if it's a :
	// - Redirection back from Gdrive with the authorized token
	// - Or a new oAuth URL needs to be provided

	stateToken := c.Query("state")
	if stateToken != "" {
		// Getting new token from Gdrive
		tokenCode := c.Query("code")
		if stateToken != "state-token" || tokenCode == "" {
			c.SecureJSON(http.StatusInternalServerError, gin.H{"error": "auth_error", "description": "Invalid token"})
			return
		}
		os.Stdout.WriteString("!! Creating google drive client with token:" + "\n" + tokenCode + "\n")
		Gds.CreateClient(c, tokenCode)
		c.SecureJSON(http.StatusOK, gin.H{"success": "auth_success", "description": "Authentication success"})
		return
	} else {
		// First time auth. Provide auth URL to the user
		authURL, err := Gds.GetAuthURL()
		if err != nil {
			c.SecureJSON(http.StatusInternalServerError, gin.H{"error": "auth_error", "description": "Auth process error"})
			return
		} else {
			c.Redirect(http.StatusSeeOther, authURL)
			return
		}
	}
}

// SearchInDoc ...
func SearchInDoc(c *gin.Context) {

	if Gds.HasClient() == true {
		fileID := strings.TrimSpace(c.Param("id"))
		if fileID == "" {
			c.SecureJSON(http.StatusBadRequest, gin.H{"error": "id_error"})
			return
		}
		word := strings.TrimSpace(c.Query("word"))
		if word == "" {
			c.SecureJSON(http.StatusBadRequest, gin.H{"error": "word_error"})
			return
		}

		os.Stdout.WriteString("!! Searching for word on google drive file:" + "\n" + fileID + "\n" + word + "\n")
		found, err := Gds.SearchInDoc(fileID, word)
		if err != nil {
			c.SecureJSON(http.StatusInternalServerError, gin.H{"error": "search_error", "description": "Error ocurred while performing the search. File id may be invalid"})
			return
		}

		if found == false {
			c.SecureJSON(http.StatusNotFound, gin.H{"result": "word not found"})
			return
		} else {
			c.SecureJSON(http.StatusOK, gin.H{"result": "word found"})
			return
		}
	} else {
		Auth(c)
	}
}

// CreateFile ...
func CreateFile(c *gin.Context) {

	if Gds.HasClient() == true {
		file := &models.File{}
		if err := c.BindJSON(file); c.Request.ContentLength == 0 || err != nil {
			c.SecureJSON(http.StatusBadRequest, gin.H{"error": "bind_error", "description": err.Error()})
			return
		}
		os.Stdout.WriteString("!! Saving file to google drive:" + "\n" + file.Name + "\n" + file.Description + "\n")
		driveFile, err := Gds.CreateFile(file)
		if err != nil {
			c.SecureJSON(http.StatusBadRequest, gin.H{"error": "save_error", "description": err.Error(), "drivefile": driveFile})
			return
		} else {
			c.SecureJSON(http.StatusOK, driveFile)
			return
		}
	} else {
		Auth(c)
	}
}
