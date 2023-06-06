package dto

type CategoryForm struct {
	Name string `json:"name" binding:"required,min=1"`
}

type CategoryDto struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
