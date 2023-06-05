package services

import (
	"online.shop.autmaple.com/internal/configs/db"
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/models"
)

func InsertBrand(brandForm *dto.BrandForm) error {
	tx, err := db.GetMysqlDB().Begin()
	if err != nil {
		return err
	}
	brand := *&models.Brand{Name: brandForm.Name, Image: brandForm.Image}
	err = brand.Insert(tx)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func QueryBrand(id int) (*dto.BrandDto, error) {
	brand := &models.Brand{ID: id}
	err := brand.QueryById(nil)
	if err != nil {
		return nil, err
	}
	brandDto := &dto.BrandDto{
		Name:  brand.Name,
		Image: brand.Image,
	}
	return brandDto, nil
}
func PageQueryBrand(offset, size int) ([]*dto.BrandDto, error) {
	brand := &models.Brand{}
	brandList, err := brand.PageQuery(nil, offset, size)
	if err != nil {
		return nil, err
	}
	var brandDtoList []*dto.BrandDto
	for _, brand := range brandList {
		brandDto := &dto.BrandDto{
			Name:  brand.Name,
			Image: brand.Image,
		}
		brandDtoList = append(brandDtoList, brandDto)
	}
	return brandDtoList, nil
}

func UpdateBrand(id int, brandForm *dto.BrandForm) error {
	tx, err := db.GetMysqlDB().Begin()
	brand := *&models.Brand{
		ID:    id,
		Name:  brandForm.Name,
		Image: brandForm.Image,
	}
	if err != nil {
		return err
	}
	err = brand.Update(tx)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func DeleteBrand(id int) error {
	tx, err := db.GetMysqlDB().Begin()
	if err != nil {
		return err
	}
	brand := &models.Brand{ID: id}
	err = brand.Delete(tx)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
