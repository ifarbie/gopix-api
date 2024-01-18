package controllers

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
func GetPhotos(c *gin.Context) {
	c.JSON(200, photo)
}