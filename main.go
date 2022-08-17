package main

import (
	"github.com/antoniocarlosmjr/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	serverGin := gin.Default()
	serverGin.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Domain Students
	serverGin.GET("/students", controllers.GetAllStudents)

	// Init Server
	serverGin.Run()
}
