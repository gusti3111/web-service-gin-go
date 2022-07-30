package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)

}

// func helloWorld(c *gin.Context) {
// 	c.JSON(http.StatusOK, "Hello, GO")

// }
func postAlbum(c *gin.Context) {
	var newAlbum album

	// panggil BindJSON untuk mnegganbungkan dengan api albums sebelumnya
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// menambahkan value dari dari newAlbums ke slice album

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// looping
	for _, v := range albums {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

func getAlbumByTitle(c *gin.Context) {
	title := c.Param("title")

	for _, a := range albums {
		if a.Title == title {
			c.JSON(http.StatusOK, a)

		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	// router.GET("/", helloWorld)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/albums/title", getAlbumByTitle)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")

}
