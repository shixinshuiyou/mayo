package session

import "github.com/gin-gonic/gin"

func GetSessionId(c *gin.Context) string {
	return c.Request.Header.Get("Session-Id")
}

func GetSessionScene(c *gin.Context) string {
	return c.Request.Header.Get("Session-Scene")
}
