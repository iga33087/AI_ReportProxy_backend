package router

import (
	//"log"
  "database/sql"
	//"fmt"
	"net/http"
  "AI-Proxy-backend/sqllib"
	"github.com/gin-gonic/gin"
)

func RegisterDeviceRoutes(rg *gin.RouterGroup,db *sql.DB) {
	deviceGroup := rg.Group("/device")
	{
		deviceGroup.GET("/", getDeviceList(db))
    deviceGroup.GET("/:id", getDeviceOne(db))
		deviceGroup.POST("/", postDeviceData(db))
		deviceGroup.PUT("/:id", updateDeviceData(db))
		deviceGroup.DELETE("/:id", delDeviceData(db))
	}
}

func getDeviceList(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    list, err := sqllib.GetDevice(db)
    if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    //fmt.Println("獲取成功:", list)
    c.JSON(http.StatusOK, gin.H{"data": list})
  }
}

func getDeviceOne(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    id := c.Param("id")
    list, err := sqllib.GetDeviceById(db,id)
    if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    //fmt.Println("獲取成功:", list)
    c.JSON(http.StatusOK, gin.H{"data": list})
  }
}

func postDeviceData(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    var body sqllib.Device
  
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
  
    id, err := sqllib.CreateDevice(db,body)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return 
    }
    //fmt.Println("新增成功，ID:", id)
    c.JSON(http.StatusOK, gin.H{"data": id})
  }
}

func updateDeviceData(db *sql.DB) gin.HandlerFunc {
  //id := c.Param("id")
  return func(c *gin.Context) {
    var body sqllib.Device
  
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
  
    uid, err := sqllib.UpdateDevice(db,body)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return 
    }
    //fmt.Println("修改成功，ID:", uid)
    c.JSON(http.StatusOK, gin.H{"data": uid})
  }
}

func delDeviceData(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    id := c.Param("id")
    uid, err := sqllib.DeleteDevice(db,id)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return 
    }
    //fmt.Println("刪除成功，ID:", uid)
    c.JSON(http.StatusOK, gin.H{"data": uid})
  }
}