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
	"gorm.io/gorm"
)

func UserRegister(c *gin.Context) {
	// MENGAMBIL INPUTAN JSON DARI USER
	var userRegisterInput app.UserRegisterInput
	if err := c.ShouldBindJSON(&userRegisterInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// VALIDASI INPUTAN PENGGUNA
	if _, err := govalidator.ValidateStruct(userRegisterInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// CEK PANJANG PASSWORD
	if len(userRegisterInput.Password) < 6 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "minimum password length is 6"})
		return
	}

	var user models.User
	// CEK APAKAH EMAIL SUDAH TERDAFTAR
	if err := database.DB.Where("email = ?", userRegisterInput.Email).First(&user).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "email address is already taken"})
		return
	}

	// CEK APAKAH USERNAME SUDAH TERDAFTAR
	if err := database.DB.Where("username = ?", userRegisterInput.Username).First(&user).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "username is already taken"})
		return
	}

	user = models.User{
		Username: userRegisterInput.Username,
		Email	: userRegisterInput.Email,
		Password: userRegisterInput.Password,
	}

	// HASH PASSWORD MENGGUNAKAN HELPER BCRYPT	
	if hashedPassword, err := helpers.HashPassword(user.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else {
		user.Password = hashedPassword
	}

	// INSERT KE DATABASE
	if err := database.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "sukses registrasi!"})
}

func UserLogin(c *gin.Context) {
	// MENGAMBIL INPUTAN JSON DARI USER
	var userLoginInput app.UserLoginInput
	if err := c.ShouldBindJSON(&userLoginInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// VALIDASI INPUTAN PENGGUNA
	if _, err := govalidator.ValidateStruct(userLoginInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// MENGAMBIL DATA USER UNTUK PENGECEKAN
	var user models.User
	// PENGECEKAN EMAIL
	if err := database.DB.Where("email = ?", userLoginInput.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Email not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	// MENCOCOKAN PASSWORD
	if err := helpers.ComparePassword(user.Password, userLoginInput.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "wrong password"})
		return
	}

	// PROSES PEMBUATAN JWT
	token, expTime, err := helpers.GenerateJWT(user.ID, user.Email, user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "failed to generate jwt"})
		return
	}
	// SET COOKIE
	c.SetCookie("token", token, int(expTime.Unix()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "login success"})
}

func UserUpdate(c *gin.Context) {
	// JIKA PARAM/ID TIDAK VALID
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "id invalid"})
		return
	}

	// APAKAH INI ADALAH USER YANG SAMA
	tokenString, _ := c.Cookie("token")
	claims, _, err := helpers.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	if userID != int(claims.ID) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized, you're not the user"})
		return
	}

	// AMBIL INPUTAN USER
	var userUpdateInput app.UserUpdateInput
	if err := c.ShouldBindJSON(&userUpdateInput); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	// VALIDASI INPUTAN USER
	if _, err := govalidator.ValidateStruct(userUpdateInput); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "please enter password"})
		return
	}

	// MENGAMBIL DATA UNTUK PENGECEKAN
	var user models.User
	// CEK EMAIL APAKAH SUDAH TERDAFTAR
	if err := database.DB.Where("email = ? AND id != ?", userUpdateInput.NewEmail, userID).First(&user).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "email address is already taken"})
		return
	}
	// CEK USERNAME APAKAH SUDAH TERDAFTAR
	if err := database.DB.Where("username = ? AND id != ?", userUpdateInput.NewUsername, userID).First(&user).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "username is not available"})
		return
	}

	// MENGAMBIL DATA USER YANG SAAT INI
	if err := database.DB.Where("id = ?", claims.ID).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "user not found"})
		return
	}
	// CEK APAKAH PASSWORD BENAR
	if err := helpers.ComparePassword(user.Password, userUpdateInput.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "wrong password!"})
		return
	}

	// UBAH JIKA USER INPUT
	if userUpdateInput.NewUsername != "" {
		user.Username = userUpdateInput.NewUsername
	}
	if userUpdateInput.NewEmail != "" {
		user.Email = userUpdateInput.NewEmail
	}
	if userUpdateInput.NewPassword != "" {
		// JIKA PASSWORD BARU KURANG DARI 6
		if len(userUpdateInput.NewPassword) < 6 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "minimum password length is 6"})
			return
		}
		// JIKA LOLOS AKAN DI HASH
		if hashedPassword, err := helpers.HashPassword(userUpdateInput.NewPassword); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			user.Password = hashedPassword
		}
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update success"})
}

func UserDelete(c *gin.Context) {
	// CEK JIKA ID TIDAK VALID
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "id invalid"})
		return
	}

	// CEK APAKAH INI ADALAH USER YANG SAMA
	tokenString, _ := c.Cookie("token")
	claims, _, err := helpers.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	if userID != int(claims.ID) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized, you're not the user"})
		return
	}

	// MENGAMBIL DATA USER SAAT INI
	var user models.User
	if err := database.DB.Where("id = ?", claims.ID).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "user not found!"})
		return
	}

	// MENGHAPUS DATA USER SAAT INI
	if err := database.DB.Delete(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "cannot delete user"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}

func UserLogout(c *gin.Context) {
	if value, err := c.Cookie("token"); err != nil || value == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "user not login yet"})
		return;
	}

	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}