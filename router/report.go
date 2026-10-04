package router

import (
	//"log"
	"fmt"
  "time"
	"database/sql"
	"net/http"
	"AI-Proxy-backend/lib"
	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(rg *gin.RouterGroup,db *sql.DB) {
	reportGroup := rg.Group("/report")
	{
		reportGroup.POST("/test", getTest(db))
		reportGroup.POST("/", postReportData(db))
	}
}

func getTest(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    var body map[string]any
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    list, err := lib.GetDeviceById(db,body["deviceId"])
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    fmt.Println("原始資料",list.UserTrafficData.Top20_UserTraffic_Group_Ranking) //Top20_UserTraffic_Ranking
    jsonData,err := lib.TextToJSONArray(list.UserTrafficData.Top20_UserTraffic_Group_Ranking) //Top20_UserTraffic_Ranking
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "JSON無法解析"})
        return
    }
    jsonC,reTable := lib.ValueConvert(jsonData)
    for jsonCIndex,jsonCVal := range jsonC {
      fmt.Println("去識別化資料",jsonCIndex," - ",jsonCVal)
    }

    for reTableIndex,reTableVal := range reTable {
      fmt.Println("對照表",reTableIndex," - ",reTableVal.RestoreTable)
    }

    c.JSON(http.StatusOK, gin.H{"data": "OK"})
  }
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

    timestamp := time.Now().Unix()
    hash := lib.HashByKey(string(timestamp) + list.UUID + list.ProductId,list.Key)
    
    header := map[string]string {
		"Authorization":hash,
	  }
      payload := map[string]any {
	  	"uuid":list.UUID,
	  	"productId":list.ProductId,
      "timestamp":timestamp,
	  }
    
    res,err := lib.Call("http://localhost:8090","GET",header,payload)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  	    return 
    }
    c.JSON(http.StatusOK, gin.H{"data": res})
  }
}