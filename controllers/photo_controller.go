package controllers

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/app"
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/database"
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/helpers"
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/models"
)

func UserUploadPhoto(c *gin.Context) {
	// AMBIL INPUTAN USER
	var userPhotoInput app.UserPhotoInput
	if err := c.ShouldBindJSON(&userPhotoInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// VALIDASI INPUTAN USER
	if _, err := govalidator.ValidateStruct(userPhotoInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// PENGECEKAN PHOTO_URL YANG DIINPUT
	if !helpers.IsImage(userPhotoInput.PhotoUrl) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "make sure the photo extension is jpg/jpeg/png"})
		return
	}

	// PROSES PEMBUATAN PHOTO
	// 1. MENGAMBIL USERID SAAT INI
	tokenString, _ := c.Cookie("token")
	claims, _, err := helpers.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// 2. MEMBUAT PHOTO
	// CEK APAKAH PHOTO URL SAMA
	var photo models.Photo
	if err := database.DB.Where("photo_url = ? AND user_id = ?", userPhotoInput.PhotoUrl, claims.ID).First(&photo).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "please choose different photo"})
		return
	}

	if userPhotoInput.Caption == "" {
		userPhotoInput.Caption = "user not set photo's caption"
	}
	photo = models.Photo{
		Title: userPhotoInput.Title,
		Caption: userPhotoInput.Caption,
		PhotoUrl: userPhotoInput.PhotoUrl,
		UserID: claims.ID,
	}

	if err := database.DB.Create(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "photo uploaded"})
}

func UserGetPhotos(c *gin.Context) {
	// MENGAMBIL ID USER SAAT INI
	tokenString, _ := c.Cookie("token")
	claims, _, err := helpers.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}	
	
	// MENGAMBIL DATA
	var userGetAllphotos []app.UserGetAllPhotos
	if err := database.DB.Table("photos").Select("photos.id, photos.user_id, photos.title, photos.caption, photos.photo_url, photos.created_at, photos.updated_at").Joins("JOIN users ON users.id = photos.user_id").Where("photos.user_id = ?", claims.ID).Order("photos.id desc").Find(&userGetAllphotos).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"photos": userGetAllphotos})
}

func UserUpdatePhoto(c *gin.Context) {
	// CEK ID ENDPOINT
	photoID, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id invalid!"})
		return
	}

	// AMBIL INPUTAN USER
	var userUpdatePhotoInput app.UserUpdatePhotoInput
	if err := c.ShouldBindJSON(&userUpdatePhotoInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// VALIDASI INPUTAN USER
	if _, err := govalidator.ValidateStruct(userUpdatePhotoInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// MENGAMBIL DATA PHOTO YANG SAAT INI
	var photo models.Photo
	if err := database.DB.Table("photos").Where("id = ?", photoID).First(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "photo not found"})
		return
	}

	// AMBIL TOKEN DAN ID USER SAAT INI
	tokenString, _ := c.Cookie("token")
	claims, _, err := helpers.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	
	// PERIZINAN
	if photo.UserID != claims.ID {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "you have no access to edit this photo"})
		return
	}

	// AMBIL DATA USER SAAT INI UNTUK COMPARE PASSWORD
	var currentUser models.User
	if err := database.DB.Table("users").Where("id = ?", claims.ID).First(&currentUser).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := helpers.ComparePassword(currentUser.Password, userUpdatePhotoInput.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "wrong password"})
		return
	}

	// PENGECEKAN INPUTAN USER
	var tempPhoto models.Photo
	if userUpdatePhotoInput.NewPhotoUrl != "" {
		if !helpers.IsImage(userUpdatePhotoInput.NewPhotoUrl) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "make sure the photo extension is jpg/jpeg/png"})
			return
		}
		if err := database.DB.Where("photo_url = ? AND user_id = ? AND id != ?", userUpdatePhotoInput.NewPhotoUrl, claims.ID, photoID).First(&tempPhoto).Error; err == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "please choose different photo"})
			return
		}
		photo.PhotoUrl = userUpdatePhotoInput.NewPhotoUrl
	}
	if userUpdatePhotoInput.NewTitle != "" {
		photo.Title = userUpdatePhotoInput.NewTitle
	}
	if userUpdatePhotoInput.NewCaption != "" {
		photo.Caption = userUpdatePhotoInput.NewCaption
	}

	if err := database.DB.Save(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "your photo succesfully changed",
	})
}

func UserDeletePhoto(c *gin.Context) {
	// CEK ID ENDPOINT
	photoID, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id invalid!"})
		return
	}

	// MENGAMBIL TOKEN USER SAAT INI
	tokenString, _ := c.Cookie("token")
	claims, _, err := helpers.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	// MENGAMBIL DATA PHOTO YANG SAAT INI
	var currentPhoto models.Photo
	if err := database.DB.Where("id = ? AND user_id = ?", photoID, claims.ID).First(&currentPhoto).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "photo not found"})
		return
	}

	// PERIZINAN
	if currentPhoto.UserID != claims.ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "you have no access to do that"})
		return
	}

	// HAPUS FOTO
	if err := database.DB.Delete(&currentPhoto).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete photo success",
	})
}