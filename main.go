package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album 代表專輯資料的結構體
type Album struct {
	ID     string  `json:"id" binding:"required"`
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// 模擬記憶體資料庫
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
}

func main() {
	// 1. 初始化預設的 Gin 引擎（內建 Logger 和 Recovery 中間件）
	router := gin.Default()

	// 2. 路由群組（適用於版本控制）
	v1 := router.Group("/")
	{
		// GET: 取得所有專輯列表
		v1.GET("/albums", getAlbums)

		// GET: 透過路徑參數（Param）取得特定專輯
		v1.GET("/albums/:id", getAlbumByID)

		// POST: 新增專輯（綁定 JSON 請求體）
		v1.POST("/albums", postAlbums)
	}

	// 3. 啟動伺服器，預設監聽 0.0.0.0:8080
	router.Run(":8080")
}

// getAlbums 回傳所有專輯的 JSON 資料
func getAlbums(c *gin.Context) {
	// 使用 Context.JSON 序列化結構體並回應 200 OK
	c.JSON(http.StatusOK, albums)
}

// getAlbumByID 示範如何獲取 URL 路徑參數 (:id)
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// 也可以獲取查詢參數（例如：/albums/1?search=rock）
	// searchQuery := c.Query("search")

	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums 示範如何解析並自動驗證前端傳來的 JSON 請求體
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// BindJSON 會根據結構體的 tag (binding:"required") 進行自動驗證
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}