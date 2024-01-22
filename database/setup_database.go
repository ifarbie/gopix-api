package database

import (
	"github.com/ifarbie/task-5-pbi-btpns-fariz-rifky-berliano/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(("root:@tcp(localhost:3306)/task_5_pbi_btpns_fariz?parseTime=true&loc=Local")))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.User{}, &models.Photo{})

	DB = database
}