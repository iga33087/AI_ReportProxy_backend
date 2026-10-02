package main

import (
	"net/http"
    //"AI-Proxy-backend/lib"
	"AI-Proxy-backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/")
	{
		v1.GET("/", getTest)
	}
	router.RegisterDeviceRoutes(v1)
	router.RegisterReportRoutes(v1)
	r.Run(":8080")
}

func getTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "HI"})
}