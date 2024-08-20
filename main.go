// main.go
package main

import (
	"ponderadas3m9/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/greeting", handlers.GetGreeting)
	router.Run(":8080")
}
