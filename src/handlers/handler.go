// handlers/handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGreeting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
