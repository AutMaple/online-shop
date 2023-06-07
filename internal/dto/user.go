package dto

type UserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserForm struct {
	Name  string `json:"name" binding:"required,min=1"`
	Email string `json:"email" binding:"required,min=1"`
	Phone string `json:"phone" binding:"required,min=1"`
}
