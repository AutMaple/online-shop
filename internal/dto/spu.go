package dto

import "online.shop.autmaple.com/internal/models"

type SpuDto struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Brand    *models.Brand    `json:"brand"`
	Category *models.Category `json:"category"`
	Attrs    []*AttrDto       `json:"attrs"`
}

type AttrDto struct {
	ID      int          `json:"id"`
	Attr    string       `json:"attr"`
	Options []*OptionDto `json:"options"`
}

type OptionDto struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type SpuForm struct {
	Name     string      `json:"name" binding:"required,min=1"`
	Brand    int         `json:"brand" binding:"required,min=1"`
	Category int         `json:"category" binding:"required,min=1"`
	Attrs    []*AttrForm `json:"attrs" binding:"required,dive,min=1"`
}

type AttrForm struct {
	Attr    string   `json:"attr" binding:"required,min=1"`
	Options []string `json:"options" binding:"required,min=1"`
}
