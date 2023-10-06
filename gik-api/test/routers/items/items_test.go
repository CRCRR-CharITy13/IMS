package items

import (
	"GIK_Web/database"
	"GIK_Web/env"
	"GIK_Web/src/routers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// global variables to represent the test environment
var Database *gorm.DB
var TestServer *gin.Engine

// ---------------------------------------------------
func ConnectDatabase() {
	env.IsLocalDB = true
	// create the test database in the same folder
	env.SqliteURI = "./test-gik-ims-localdb.sqlite"
	env.SkipMigrations = false
	database.ConnectDatabase()
	// TODO: add testing function here

	// Note: by the end, make sure to clear table

}

func InitRouter() {
	TestServer = routers.InitRouter()
}

type ConnectServerResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func TestConnectServer(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	TestServer.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response ConnectServerResponse
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)
	assert.True(t, response.Success)
	assert.Equal(t, "Pong!", response.Message)
}

func TestCreateFirstAdmin(t *testing.T) {

	//TestServer.GET("/auth/first_admin", auth.CreateFirstAdmin)
	req, _ := http.NewRequest("GET", "/auth/first_admin", nil)
	query := req.URL.Query()
	query.Add("password", "password")
	req.URL.RawQuery = query.Encode()
	w := httptest.NewRecorder()
	TestServer.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMain(m *testing.M) {
	ConnectDatabase()
	InitRouter()
	m.Run()
}
