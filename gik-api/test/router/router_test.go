package routers

import (
	"GIK_Web/src/routers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type initRouterResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func TestInitRouter(t *testing.T) {
	router := routers.InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response initRouterResponse
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)
	assert.True(t, response.Success)
	assert.Equal(t, "Pong!", response.Message)
}
