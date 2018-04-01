package items

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
		{"GET", "/item/1", 200, "{id:1}"},
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.Run("80")
	router.Any("/", ProcessResponse)

	for _, testCase := range testCases {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(testCase.method, testCase.url, nil)
		router.ServeHTTP(w, r)
	}
}
