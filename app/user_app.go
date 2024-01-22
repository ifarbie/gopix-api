package app

type UserRegisterInput struct {
	Username        string `valid:"required" json:"username"`
	Email           string `valid:"required, email" json:"email"`
	Password        string `valid:"required" json:"password"`
}

type UserLoginInput struct {
	Email           string `valid:"required, email" json:"email"`
	Password        string `valid:"required" json:"password"`
}

type UserUpdateInput struct {
	NewUsername string `json:"new_username"`
	NewEmail    string `valid:"email" json:"new_email"`
	NewPassword string `json:"new_password"`
	Password	string `valid:"required" json:"password"`
}