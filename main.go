package main

import (
	"log"
	"net/http"
    "AI-Proxy-backend/lib"
	"AI-Proxy-backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := lib.InitDB()
	if err != nil {
		log.Fatalf("資料庫連線失敗: %v", err)
	}
	defer db.Close()

	r := gin.Default()
	v1 := r.Group("/")
	{
		v1.GET("/", getTest)
	}
	router.RegisterDeviceRoutes(v1,db)
	router.RegisterReportRoutes(v1,db)
	r.Run(":8080")
}

func getTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "HI"})
}