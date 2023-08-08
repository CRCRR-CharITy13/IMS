/*
login finalizes the login process

Contains the code to run the login process.
*/

package auth

import (
	"GIK_Web/database"
	"GIK_Web/env"
	"GIK_Web/types"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// Data structure for request
type loginRequest struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	VerificationJWT string `json:"verificationJWT" binding:"required"`
}

func Login(c *gin.Context) {
	// Get query of remember me status
	rememberMe := c.Query("remember")

	// Get login request JWT
	json := loginRequest{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	// Attempt to parse tthe JWT
	token, err := jwt.Parse(json.VerificationJWT, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.JWTSigningPassword), nil
	})
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid verification",
		})
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	// check if JWT is valid
	if !token.Valid {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid verification",
		})
		return
	}

	// check if JWT is expired
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Verification is expired",
		})
		return
	}

	// check if username matches up
	json.Username = strings.ToLower(json.Username)

	if claims["username"].(string) != json.Username {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Verification does not match",
		})
		return
	}

	// check if user exists
	user := types.User{}
	if err := database.Database.Model(&types.User{}).Where(&types.User{
		Username: json.Username,
	}).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid username or password",
			"code":    1,
		})
		return
	}

	// create session
	var days int64

	rememberMeValue, err := strconv.ParseBool(rememberMe)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	if rememberMeValue {
		days = 7
	} else {
		days = 1
	}

	session := types.Session{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		CreatedAt: time.Now().Unix(),
		ExpiresAt: time.Now().Unix() + (60 * 60 * 24 * days),
	}

	if err := database.Database.Create(&session).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to create session",
		})
		return
	}

	session_cookie := sessions.Default(c)
	session_cookie.Set("session", session.ID)
	session_cookie.Options(sessions.Options{
		Path:   "/",
		MaxAge: 60 * 60 * 24 * int(days),
	})

	if err := session_cookie.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Login successful",
		"data":    session.ID,
	})

}
