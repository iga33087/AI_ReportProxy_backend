package main

import (
	"net/http"
    //"AI-Proxy-backend/lib"
	"AI-Proxy-backend/router"
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
	router := gin.Default()
	v1 := router.Group("/")
	{
		v1.GET("/albums", getAlbums)
		v1.GET("/albums/:id", getAlbumByID)
		v1.POST("/albums", postAlbums)
	}
	device.RegisterRoutes(v1)
	router.Run(":8080")
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}