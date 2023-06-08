package brand

import (
	"online.shop.autmaple.com/internal/models"
)

type Form struct {
	Name  string `json:"name" binding:"required,min=1"`
	Image string `json:"image" binding:"required,min=1"`
}

type Dto struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func InsertBrand(brandForm *Form) error {
	brand := *&models.Brand{Name: brandForm.Name, Image: brandForm.Image}
	err := brand.Insert(nil)
	if err != nil {
		return err
	}
	return nil
}

func QueryBrand(id int) (*Dto, error) {
	brand := &models.Brand{ID: id}
	err := brand.QueryById(nil)
	if err != nil {
		return nil, err
	}
	brandDto := &Dto{
		Name:  brand.Name,
		Image: brand.Image,
	}
	return brandDto, nil
}
func PageQueryBrand(offset, size int) ([]*Dto, error) {
	brand := &models.Brand{}
	brandList, err := brand.PageQuery(nil, offset, size)
	if err != nil {
		return nil, err
	}
	var brandist []*Dto
	for _, brand := range brandList {
		brand := &Dto{
			Name:  brand.Name,
			Image: brand.Image,
		}
		brandist = append(brandist, brand)
	}
	return brandist, nil
}

func UpdateBrand(id int, brandForm *Form) error {
	brand := *&models.Brand{
		ID:    id,
		Name:  brandForm.Name,
		Image: brandForm.Image,
	}
	err := brand.Update(nil)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBrand(id int) error {
	brand := &models.Brand{ID: id}
	err := brand.Delete(nil)
	if err != nil {
		return err
	}
	return nil
}
