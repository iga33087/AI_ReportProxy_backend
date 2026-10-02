package router

import (
	//"log"
	//"fmt"
	"net/http"
	"AI-Proxy-backend/lib"
	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(rg *gin.RouterGroup) {
	reportGroup := rg.Group("/report")
	{
		reportGroup.GET("/", getTest)
		reportGroup.POST("/", postReportData)
	}
}

func getTest(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"data": "1111"})
}

func postReportData(c *gin.Context) {
  header := map[string]string {}
  payload := map[string]any {}
  res,err := lib.Call("https://example.org/","GET",header,payload)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return 
  }
  c.JSON(http.StatusOK, gin.H{"data": res})
}