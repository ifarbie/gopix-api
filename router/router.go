package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")
	users.GET("/photos", controllers.GetPhotos)

	return r
}