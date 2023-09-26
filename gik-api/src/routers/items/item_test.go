package items

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestListItem(t *testing.T) {
	r := gin.Default()
	r.GET("/items/list", ListItem)

	// create a test HTTP response recorder
	w := httptest.NewRecorder()

	//create a test HTTP request
	req, err := http.NewRequest("GET", "items/list", nil)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	r.ServeHTTP(w, req)

	//check the response status code

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}
