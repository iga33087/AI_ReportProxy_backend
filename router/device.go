package router

import (
	//"log"
	"fmt"
	"net/http"
	"AI-Proxy-backend/lib"
	"github.com/gin-gonic/gin"
)

func RegisterDeviceRoutes(rg *gin.RouterGroup) {
	deviceGroup := rg.Group("/device")
	{
		deviceGroup.GET("/", getList)
		deviceGroup.POST("/", postData)
		deviceGroup.PUT("/:id", updateData)
		deviceGroup.DELETE("/:id", delData)
	}
}

func getList(c *gin.Context) {
  db, err := lib.InitDB()
  if err != nil {
  	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
  defer db.Close()
  list, err := lib.GetDevice(db)
  if err != nil {
  	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
  fmt.Println("獲取成功:", list)
  c.JSON(http.StatusOK, gin.H{"data": list})
}

func postData(c *gin.Context) {
  var body lib.Device

  if err := c.ShouldBindJSON(&body); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{
          "error": err.Error(),
      })
      return
  }

  db, err := lib.InitDB()
  if err != nil {
  	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
  defer db.Close()
  id, err := lib.CreateDevice(db,body)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return 
  }
  fmt.Println("新增成功，ID:", id)
  c.JSON(http.StatusOK, gin.H{"data": id})
}

func updateData(c *gin.Context) {
  //id := c.Param("id")
  var body lib.Device

  if err := c.ShouldBindJSON(&body); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
  }

  db, err := lib.InitDB()
  if err != nil {
  	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
  defer db.Close()
  uid, err := lib.UpdateDevice(db,body)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return 
  }
  fmt.Println("修改成功，ID:", uid)
  c.JSON(http.StatusOK, gin.H{"data": uid})

}

func delData(c *gin.Context) {
  id := c.Param("id")

  db, err := lib.InitDB()
  if err != nil {
  	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
  defer db.Close()
  uid, err := lib.DeleteDevice(db,id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return 
  }
  fmt.Println("刪除成功，ID:", uid)
  c.JSON(http.StatusOK, gin.H{"data": uid})
}