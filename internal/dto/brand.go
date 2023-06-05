package dto

type BrandForm struct {
	Name  string `json:"name" binding:"required,min=1"`
	Image string `json:"image" binding:"required,min=1"`
}

type BrandDto struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
