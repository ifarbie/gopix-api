package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/controllers"
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/database"
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	database.ConnectDatabase()

	usersPublic := r.Group("/users")
	usersPublic.POST("/register", controllers.UserRegister)
	usersPublic.POST("/login", controllers.UserLogin)
	
	usersProtected := r.Group("/users")
	usersProtected.Use(middlewares.RequireAuth())
	usersProtected.GET("/logout", controllers.UserLogout)
	usersProtected.PUT("/:userId", controllers.UserUpdate)
	usersProtected.DELETE("/:userId", controllers.UserDelete)

	photosProtected := r.Group("/photos")
	photosProtected.Use(middlewares.RequireAuth())
	photosProtected.POST("", controllers.UserUploadPhoto)
	photosProtected.GET("", controllers.UserGetPhotos)
	photosProtected.PUT("/:photoId", controllers.UserUpdatePhoto)
	// photosProtected.DELETE("/:photoid", controllers.UserDeletePhoto)

	return r
}