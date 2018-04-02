package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Test ...
func Test(t *testing.T) {

	testCases := []struct {
		method       string
		url          string
		expectedCode int
		expectedBody string
	}{
		{"GET", "/item/1", http.StatusOK, "{id:1}"},
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Any("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.Run("80")

	for _, testCase := range testCases {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(testCase.method, testCase.url, nil)
		router.ServeHTTP(w, r)
	}
}
