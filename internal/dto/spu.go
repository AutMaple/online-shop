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

type SpuForm struct {
	Name     string      `json:"name" binding:"required"`
	Brand    int         `json:"brand" binding:"required"`
	Category int         `json:"category" binding:"required"`
	Attrs    []*AttrForm `json:"attrs" binding:"required"`
}

type AttrForm struct {
	Attr    string   `json:"attr" binding:"required"`
	Options []string `json:"options" binding:"required"`
}

type OptionDto struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}
