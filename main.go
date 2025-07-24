package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type album struct {
	ID string `json: "id"`
	Title string `json: "title"`
	Artist string `json: "artist"`
	Price float64 `json: "price"`
}

var albums = []album{
	{ID: "1", Title: "Blue", Artist: "young kai", Price: 56.99},
	{ID: "2", Title: "Blue2", Artist: "young kai", Price: 54.99},
	{ID: "1", Title: "Blue3", Artist: "young kai", Price: 51.99},

}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}