package dto

type SpuDto struct {
	Name     string    `json:"name"`
	Brand    int       `json:"brand"`
	Category int       `json:"category"`
	Attrs    []AttrDto `json:"attrs"`
}

type AttrDto struct {
	Attr    string   `json:"attr"`
	Options []string `json:"options"`
}
