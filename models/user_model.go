package models

import "time"

type User struct {
	ID       	uint		`gorm:"primaryKey" json:"id"` 
	Username 	string		`gorm:"not null" json:"username"`
	Email    	string		`gorm:"not null;unique" json:"email"`
	Password 	string		`gorm:"not null" json:"password"`
	Photos 		[]Photo		`gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"photos"`
	CreatedAt 	time.Time	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 	time.Time	`gorm:"autoUpdateTime" json:"updated_at"`
}