package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	// End session

	// Reset cookie
	session_cookie := sessions.Default(c)
	session_cookie.Set("session", "")
	session_cookie.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	if err := session_cookie.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Logout successful",
	})
}
