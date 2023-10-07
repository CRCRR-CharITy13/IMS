package items

import (
	"GIK_Web/database"
	"GIK_Web/env"
	"GIK_Web/src/routers"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// global variables to represent the test environment
var Database *gorm.DB
var TestRouter *gin.Engine
var tokenString string
var sessionID string
var cookies string

// ---------------------------------------------------
func ConnectDatabase() {
	env.IsLocalDB = true
	// create the test database in the same folder
	env.SqliteURI = "./test-gik-ims-localdb.sqlite"
	env.SkipMigrations = false
	database.ConnectDatabase()
	// Attention: by the end, make sure to clear the tables

}

func InitRouter() {
	TestRouter = routers.InitRouter()
}

type ConnectServerResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func TestConnectServer(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	TestRouter.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response ConnectServerResponse
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)
	assert.True(t, response.Success)
	assert.Equal(t, "Pong!", response.Message)
}

func TestCreateFirstAdmin(t *testing.T) {
	//TestRouter.GET("/auth/first_admin", "auth/first_admin")
	req, _ := http.NewRequest("GET", "/auth/first_admin", nil)
	query := req.URL.Query()
	query.Add("password", "password")
	req.URL.RawQuery = query.Encode()
	w := httptest.NewRecorder()
	TestRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListItem_Without_Login(t *testing.T) {
	req, _ := http.NewRequest("GET", "/items/list", nil)
	w := httptest.NewRecorder()
	TestRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

type PreLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PreLoginResponse struct {
	Success bool
	Message string
	Data    string
}

func TestPreLogin(t *testing.T) {
	a := assert.New(t)
	//
	user := PreLoginRequest{
		Username: "admin",
		Password: "password",
	}
	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}
	//
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/prelogin", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	TestRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	preLoginRes := PreLoginResponse{}
	err = json.Unmarshal(w.Body.Bytes(), &preLoginRes)
	tokenString = preLoginRes.Data
	//
}

type LoginRequest struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	VerificationJWT string `json:"verificationJWT" binding:"required"`
}
type LoginResponse struct {
	Success bool
	Message string
	Data    string
}

func TestLogin(t *testing.T) {
	a := assert.New(t)
	//
	// fmt.Printf("tokenString = %s", tokenString)

	user := LoginRequest{
		Username:        "admin",
		Password:        "password",
		VerificationJWT: tokenString,
	}
	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}
	//

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	query := req.URL.Query()
	query.Add("remember", "true")
	req.URL.RawQuery = query.Encode()
	TestRouter.ServeHTTP(w, req)
	cookies = w.HeaderMap["Set-Cookie"][0]
	fmt.Print(cookies)
	assert.Equal(t, http.StatusOK, w.Code)

	loginResponse := LoginResponse{}
	err = json.Unmarshal(w.Body.Bytes(), &loginResponse)
	sessionID = loginResponse.Data
	fmt.Printf("TestLogin, SessionID = %v\n", sessionID)
}

type ListItemResponse struct {
	Success bool
	Message string
	Data    string
}
type AddNewItemRequest struct {
	SKU        string  `json:"sku" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Size       string  `json:"size" binding:"required"`
	Price      float32 `json:"price" binding:"required"`
	StockTotal int     `json:"quantity" binding:"required"`
}

func TestAddItem_With_Login(t *testing.T) {
	a := assert.New(t)

	item := AddNewItemRequest{
		SKU:        "TESTSKU02",
		Name:       "password",
		Size:       "M",
		Price:      10.0,
		StockTotal: 2,
	}
	reqBody, err := json.Marshal(item)
	if err != nil {
		a.Error(err)
	}
	//

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/items/add", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookies)
	query := req.URL.Query()
	req.URL.RawQuery = query.Encode()
	TestRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestListItem_With_Login(t *testing.T) {
	req, _ := http.NewRequest("GET", "/items/list", nil)
	req.Header.Set("Cookie", cookies)
	w := httptest.NewRecorder()
	TestRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Print(w.Body)
}
func TestMain(m *testing.M) {
	ConnectDatabase()
	InitRouter()
	m.Run()
}
