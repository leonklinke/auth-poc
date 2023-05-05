package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SessionAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorized, err := authorizeSession(c)
		if !authorized {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
	}
}
