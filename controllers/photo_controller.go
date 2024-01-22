package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/database"
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/models"
)

func GetPhotos(c *gin.Context) {
	var photos []models.Photo

	database.DB.Find(&photos)

	c.JSON(http.StatusOK, gin.H{"photos": photos})
}