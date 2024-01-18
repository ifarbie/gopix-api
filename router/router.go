package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/models"
)

var photo = models.Photo{
	ID: 1,
	Title: "Photo", 
	Caption: "Caption",
	PhotoUrl: "img.com",
	UserID: 2,
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/photos", func(c *gin.Context) {
		c.JSON(200, photo)
	})

	return r
}