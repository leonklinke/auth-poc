package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

const BEARER_SCHEMA = "Bearer"

func getSessionToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	return tokenString
}

func authorizeSession(c *gin.Context) (bool, error) {
	token := getSessionToken(c)
	path := c.Request.URL.Path
	log.Println("token", token)
	log.Println("path", path)
}
