package services

import (
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/models"
)

func QueryCategory(id int) (*dto.CategoryDto, error) {
	c := models.Category{ID: id}
	err := c.QueryById(nil)
	if err != nil {
		return nil, err
	}
	categoryDto := &dto.CategoryDto{
		ID:   c.ID,
		Name: c.Name,
	}
	return categoryDto, nil
}

func PageQueryCategory(offset, size int) ([]*dto.CategoryDto, error) {
	c := models.Category{}
	categoryList, err := c.PageQuery(nil, offset, size)
	if err != nil {
		return nil, err
	}
	var categoryDtoList []*dto.CategoryDto
	for _, category := range categoryList {
		categoryDto := &dto.CategoryDto{
			ID:   category.ID,
			Name: category.Name,
		}
		categoryDtoList = append(categoryDtoList, categoryDto)
	}
	return categoryDtoList, nil
}

func UpdateCategory(id int, categoryForm *dto.CategoryForm) error {
	c := models.Category{
		ID:   id,
		Name: categoryForm.Name,
	}
	err := c.Update(nil)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(id int) error {
	c := models.Category{
		ID: id,
	}
	err := c.Delete(nil)
	if err != nil {
		return err
	}
	return nil
}

func InsertCategory(categoryForm *dto.CategoryForm) error {
	c := models.Category{
		Name: categoryForm.Name,
	}
	err := c.Insert(nil)
	if err != nil {
		return err
	}
	return nil
}
