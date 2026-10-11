package router

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Roc() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := map[string]bool{
			"http://localhost:5173": true,
		}

		origin := c.Request.Header.Get("Origin")

		if allowedOrigins[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods",
				"POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers",
				"Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Max-Age", "86400")
			c.Header("Vary", "Origin")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// Public endpoints that do not require authentication.
func getWhite() map[string]bool {
	return map[string]bool{
		"/test/list": true,
	}
}

// Authentication placeholder for the test project.
// Real JWT validation can be added later.
func Filter() gin.HandlerFunc {
	whitelist := getWhite()

	return func(c *gin.Context) {
		if c.Request.Method != http.MethodPost &&
			c.Request.Method != http.MethodGet {
			c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
				"success": false,
				"message": "Method not allowed",
			})
			return
		}

		if whitelist[c.Request.URL.Path] {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Authentication not implemented",
		})
	}
}

func Recovery(c *gin.Context) {
	defer func() {
		if recovered := recover(); recovered != nil {
			debug.PrintStack()

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Internal server error",
			})
		}
	}()

	c.Next()
}
