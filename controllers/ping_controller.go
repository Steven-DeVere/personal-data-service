package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a function that returns a "pong". Useful when verifying that the server is running
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
