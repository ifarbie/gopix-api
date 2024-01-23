package app

import "time"

type UserPhotoInput struct {
	Title    string `json:"title" valid:"required"`
	Caption  string `json:"caption" valid:"optional"`
	PhotoUrl string `json:"photo_url" valid:"required"`
}

type UserGetAllPhotos struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUpdatePhotoInput struct {
	NewTitle    string `json:"new_title" valid:"optional"`
	NewCaption  string `json:"new_caption" valid:"optional"`
	NewPhotoUrl string `json:"new_photo_url" valid:"optional"`
	Password 	string `json:"password" valid:"required"`
}