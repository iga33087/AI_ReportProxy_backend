package device

import (
	"log"
	"fmt"
	"net/http"
	"AI-Proxy-backend/lib"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	deviceGroup := rg.Group("/device")
	{
		deviceGroup.GET("/", GetList)
		deviceGroup.POST("/", PostData)
		deviceGroup.PUT("/:id", UpdateData)
		deviceGroup.DELETE("/:id", DelData)
	}
}

func GetList(c *gin.Context) {
  db, err := sqlite.InitDB()
  if err != nil {
  	log.Fatal(err)
  }
  defer db.Close()
  list, err := sqlite.GetDevice(db)
  if err != nil {
  	log.Fatal(err)
  }
  fmt.Println("獲取成功:", list)
  c.JSON(http.StatusOK, gin.H{"data": list})
}

func PostData(c *gin.Context) {
  var body sqlite.Device

  if err := c.ShouldBindJSON(&body); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{
          "error": err.Error(),
      })
      return
  }

  db, err := sqlite.InitDB()
  if err != nil {
  	log.Fatal(err)
  }
  defer db.Close()
  id, err := sqlite.CreateDevice(db,body)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return 
  }
  fmt.Println("新增成功，ID:", id)
  c.JSON(http.StatusOK, gin.H{"data": id})
}

func UpdateData(c *gin.Context) {
  //id := c.Param("id")
  var body sqlite.Device

  if err := c.ShouldBindJSON(&body); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
  }

  db, err := sqlite.InitDB()
  if err != nil {
  	log.Fatal(err)
  }
  defer db.Close()
  uid, err := sqlite.UpdateDevice(db,body)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return 
  }
  fmt.Println("修改成功，ID:", uid)
  c.JSON(http.StatusOK, gin.H{"data": uid})

}

func DelData(c *gin.Context) {
  id := c.Param("id")

  db, err := sqlite.InitDB()
  if err != nil {
  	log.Fatal(err)
  }
  defer db.Close()
  uid, err := sqlite.DeleteDevice(db,id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return 
  }
  fmt.Println("刪除成功，ID:", uid)
  c.JSON(http.StatusOK, gin.H{"data": uid})
}