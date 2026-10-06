package router

import (
	//"log"
	"fmt"
  "time"
	"database/sql"
	"net/http"
	"encoding/json"
	"AI-Proxy-backend/lib"
	"AI-Proxy-backend/sqllib"
	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(rg *gin.RouterGroup,db *sql.DB) {
	reportGroup := rg.Group("/report")
	{
		reportGroup.POST("/test", getReportTest(db))
		reportGroup.GET("/", getReportList(db))
        reportGroup.GET("/:id", getReportOne(db))
		reportGroup.POST("/", postReportData(db))
		reportGroup.DELETE("/:id", delReportData(db))
	}
}

func getReportTest(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    var body map[string]any
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    list, err := sqllib.GetDeviceById(db,body["deviceId"])
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("原始資料",list)

    reportRawData,reTable := getReportRawData(list,int(body["reportType"].(float64)))
    if reportRawData == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "getReportRawData處理失敗"})
        return
    }

    /*jsonC,reTable := lib.ValueConvert(rawData.(map[string][]map[string]any))
    for jsonCIndex,jsonCVal := range jsonC {
      fmt.Println("去識別化資料",jsonCIndex," - ",jsonCVal)
    }

    for reTableIndex,reTableVal := range reTable {
      fmt.Println("對照表",reTableIndex," - ",reTableVal.RestoreTable)
    }*/
    fmt.Println("11",reportRawData,reTable)
    c.JSON(http.StatusOK, gin.H{"data": "OK"})
  }
}

func getReportList(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    list, err := sqllib.GetReport(db)
    if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    fmt.Println("獲取成功:", list)
    c.JSON(http.StatusOK, gin.H{"data": list})
  }
}

func getReportOne(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    id := c.Param("id")
    list, err := sqllib.GetReportById(db,id)
    if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    fmt.Println("獲取成功:", list)
    c.JSON(http.StatusOK, gin.H{"data": list})
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
  
    list, err := sqllib.GetDeviceById(db,body["deviceId"])
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    reportRawData,reTable := getReportRawData(list,int(body["reportType"].(float64)))
    if reportRawData == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "getReportRawData處理失敗"})
        return
    }

    //reportRawData,reTable := lib.ValueConvert(rawData.(map[string][]map[string]any))

    fmt.Println("去識別化資料",reportRawData)
    fmt.Println("對照表",reTable)

    timestamp := time.Now().Unix()
    hash := lib.HashByKey(string(timestamp) + list.UUID + list.ProductId,list.Key)
    
    header := map[string]string {
		  "Authorization":hash,
	  }
    payload := map[string]any {
	  	"uuid":list.UUID,
        "reportType":body["reportType"],
        "name":list.Name,
	  	"productId":list.ProductId,
        "timestamp":timestamp,
        "reportRawData":reportRawData,
	  }
    
    res,err := lib.Call("http://localhost:8090/call","GET",header,payload)
    if err != nil {
    	var errorJSON map[string]any
    
    	if jsonErr := json.Unmarshal([]byte(err.Error()), &errorJSON); jsonErr != nil {
    		c.JSON(http.StatusBadRequest, gin.H{
    			"error": err.Error(),
    		})
    		return
    	}
    
    	c.JSON(http.StatusBadRequest, errorJSON)
    	return
    }

	restoreRes := lib.ToRegexp(res.(map[string]any),reTable)

	restoreResJSon, err := json.Marshal(restoreRes)
	if err != nil {
		return 
	}
	reportData := sqllib.Report{
	  DeviceId:   int(body["deviceId"].(float64)),
	  ReportType: int(body["reportType"].(float64)),
	  ReportData: string(restoreResJSon),
	}
	_, createErr := sqllib.CreateReport(db,reportData)
    if createErr != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": createErr.Error()})
	    return 
    }

    c.JSON(http.StatusOK, restoreRes)
  }
}

func getReportRawData(data *sqllib.Device, reportType int) (any,map[string]*lib.ConvertItem) {
  if reportType == 1 {
      res := make(map[string]map[string]any)
		  cpuArr, err := lib.TextToJSON[map[string]any](data.HardwareData.CPU_Usage_Summary)
		  if err == nil {
		  	res["cpuUsageSummary"] = cpuArr
		  } else {
		  	res["cpuUsageSummary"] = map[string]any{}
		  }
		  memArr, err := lib.TextToJSON[map[string]any](data.HardwareData.Memory_Usage_Summary)
		  if err == nil {
		  	res["memoryUsageSummary"] = memArr
		  } else {
		  	res["memoryUsageSummary"] = map[string]any{}
		  }
		  newConnectionArr, err := lib.TextToJSON[map[string]any](data.HardwareData.New_Connection_Summary)
		  if err == nil {
		  	res["newConnectionSummary"] = newConnectionArr
		  } else {
		  	res["newConnectionSummary"] = map[string]any{}
		  }
		  activeConnectionArr, err := lib.TextToJSON[map[string]any](data.HardwareData.Active_Connection_Summary)
		  if err == nil {
		  	res["activeConnectionSummary"] = activeConnectionArr
		  } else {
		  	res["activeConnectionSummary"] = map[string]any{}
		  }
		  onlineVerificationArr, err := lib.TextToJSON[map[string]any](data.HardwareData.Online_Verification_Summary)
		  if err == nil {
		  	res["onlineVerificationSummary"] = onlineVerificationArr
		  } else {
		  	res["onlineVerificationSummary"] = map[string]any{}
		  }
		  onlineUsersArr, err := lib.TextToJSON[map[string]any](data.HardwareData.Online_Users_Summary)
		  if err == nil {
		  	res["onlineUsersSummary"] = onlineUsersArr
		  } else {
		  	res["onlineUsersSummary"] = map[string]any{}
		  }
      r,t := lib.ValueConvert2(res)
		  return r,t
  } else if reportType == 2 {
      res := make(map[string][]map[string]any)
		  rankingArr, err := lib.TextToJSON[[]map[string]any](data.UserTrafficData.Top20_UserTraffic_Ranking)
		  if err == nil {
		  	res["top20UserTrafficRanking"] = rankingArr
		  } else {
		  	res["top20UserTrafficRanking"] = []map[string]any{}
		  }
		  groupRankingArr, err := lib.TextToJSON[[]map[string]any](data.UserTrafficData.Top20_UserTraffic_Group_Ranking)
		  if err == nil {
		  	res["top20UserTrafficGroupRanking"] = groupRankingArr
		  } else {
		  	res["top20UserTrafficGroupRanking"] = []map[string]any{}
		  }
      r,t := lib.ValueConvert(res)
		  return r,t
  } else if reportType == 3 {
      res := make(map[string][]map[string]any)
		  rankingArr, err := lib.TextToJSON[[]map[string]any](data.ServiceTrafficData.Top20_ServiceTraffic_Ranking)
		  if err == nil {
		  	res["top20ServiceTrafficRanking"] = rankingArr
		  } else {
		  	res["top20ServiceTrafficRanking"] = []map[string]any{}
		  }
		  typeRankingArr, err := lib.TextToJSON[[]map[string]any](data.ServiceTrafficData.Top20_ServiceTraffic_Type_Ranking)
		  if err == nil {
		  	res["top20ServiceTrafficTypeRanking"] = typeRankingArr
		  } else {
		  	res["top20ServiceTrafficTypeRanking"] = []map[string]any{}
		  }
      r,t := lib.ValueConvert(res)
		  return r,t
  } else if reportType == 4 {
      res := make(map[string][]map[string]any)
		  rankingArr, err := lib.TextToJSON[[]map[string]any](data.DomainTrafficData.Top20_DomainTraffic_Ranking)
		  if err == nil {
		  	res["top20DomainTrafficRanking"] = rankingArr
		  } else {
		  	res["top20DomainTrafficRanking"] = []map[string]any{}
		  }
		  typeRankingArr, err := lib.TextToJSON[[]map[string]any](data.DomainTrafficData.Top20_DomainTraffic_Type_Ranking)
		  if err == nil {
		  	res["top20DomainTrafficTypeRanking"] = typeRankingArr
		  } else {
		  	res["top20DomainTrafficTypeRanking"] = []map[string]any{}
		  }
      r,t := lib.ValueConvert(res)
		  return r,t
  }
	return nil,nil
}

func delReportData(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    id := c.Param("id")
    uid, err := sqllib.DeleteReport(db,id)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return 
    }
    fmt.Println("刪除成功，ID:", uid)
    c.JSON(http.StatusOK, gin.H{"data": uid})
  }
}