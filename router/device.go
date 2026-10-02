package router

import (
	//"log"
  "database/sql"
	"fmt"
	"net/http"
	"AI-Proxy-backend/lib"
	"github.com/gin-gonic/gin"
)

func RegisterDeviceRoutes(rg *gin.RouterGroup,db *sql.DB) {
	deviceGroup := rg.Group("/device")
	{
		deviceGroup.GET("/", getList(db))
    deviceGroup.GET("/:id", getOne(db))
		deviceGroup.POST("/", postData(db))
		deviceGroup.PUT("/:id", updateData(db))
		deviceGroup.DELETE("/:id", delData(db))
	}
}

func getList(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    list, err := lib.GetDevice(db)
    if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    fmt.Println("獲取成功:", list)
    c.JSON(http.StatusOK, gin.H{"data": list})
  }
}

func getOne(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    id := c.Param("id")
    list, err := lib.GetDeviceById(db,id)
    if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    fmt.Println("獲取成功:", list)
    c.JSON(http.StatusOK, gin.H{"data": list})
  }
}

func postData(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    var body lib.Device
  
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
  
    id, err := lib.CreateDevice(db,body)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return 
    }
    fmt.Println("新增成功，ID:", id)
    c.JSON(http.StatusOK, gin.H{"data": id})
  }
}

func updateData(db *sql.DB) gin.HandlerFunc {
  //id := c.Param("id")
  return func(c *gin.Context) {
    var body lib.Device
  
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
  
    uid, err := lib.UpdateDevice(db,body)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return 
    }
    fmt.Println("修改成功，ID:", uid)
    c.JSON(http.StatusOK, gin.H{"data": uid})
  }
}

func delData(db *sql.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    id := c.Param("id")
    uid, err := lib.DeleteDevice(db,id)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return 
    }
    fmt.Println("刪除成功，ID:", uid)
    c.JSON(http.StatusOK, gin.H{"data": uid})
  }
}