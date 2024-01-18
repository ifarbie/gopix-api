package models

type User struct{
	ID uint  `gorm:"primaryKey" json:"id"`
	Username string 
	Email string
// 	Password string
// 	Photos []photo
// 	CreatedAt
// 	UpdatedAt
}