package router

import (
	//"log"
	//"fmt"
	"database/sql"
	"net/http"
	"AI-Proxy-backend/lib"
	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(rg *gin.RouterGroup,db *sql.DB) {
	reportGroup := rg.Group("/report")
	{
		reportGroup.GET("/", getTest)
		reportGroup.POST("/", postReportData(db))
	}
}

func getTest(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"data": "1111"})
}

func postReportData(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    var body map[string]any
  
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
  
    if _, ok := body["deviceId"]; !ok {
        c.JSON(http.StatusBadRequest, gin.H{"error": "沒有deviceId"})
        return
    }
  
    list, err := lib.GetDeviceById(db,body["deviceId"])
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hash := lib.HashByKey(list.UUID + body["productId"].(string),list.Key)
    
    header := map[string]string {
		"Authorization":hash,
	}
    payload := map[string]any {
		"uuid":list.UUID,
		"productId":body["productId"],
	}
    res,err := lib.Call("http://localhost:8090","GET",header,payload)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  	    return 
    }
    c.JSON(http.StatusOK, gin.H{"data": res})
  }
}