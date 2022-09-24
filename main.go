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

var albums = []album {
    {ID: "1", Title: "Colors", Artist: "Between the buried and me", Price: 56.99},
    {ID: "2", Title: "Veleno", Artist: "Fleshgod Apocalypse", Price: 59.99},
    {ID: "3", Title: "Catharsis", Artist: "YOB", Price: 39.99},
	{ID: "4", Title: "Cognitive", Artist: "Soen", Price: 56.99},
    {ID: "5", Title: "10,000 Days", Artist: "Tool", Price: 17.99},
    {ID: "6", Title: "Rays of Darkness", Artist: "MONO", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	r := gin.Default()
	r.GET("/albums", getAlbums)
	r.Run("0.0.0.0:8080")
}