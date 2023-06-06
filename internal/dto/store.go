package dto

type StoreDto struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type StoreForm struct {
	BrandIds string `json:"brandIds" binding:"required,min=1"`
	Name     string `json:"name" binding:"required,min=1"`
	Address  string `json:"address" binding:"required,min=1"`
	Phone    string `json:"phone" binding:"required,min=1"`
}
