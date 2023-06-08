package category

import (
	"online.shop.autmaple.com/internal/models"
)

type Form struct {
	Name string `json:"name" binding:"required,min=1"`
}

type Dto struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func QueryCategory(id int) (*Dto, error) {
	c := models.Category{ID: id}
	err := c.QueryById(nil)
	if err != nil {
		return nil, err
	}
	categoryDto := &Dto{
		ID:   c.ID,
		Name: c.Name,
	}
	return categoryDto, nil
}

func PageQueryCategory(offset, size int) ([]*Dto, error) {
	c := models.Category{}
	categoryList, err := c.PageQuery(nil, offset, size)
	if err != nil {
		return nil, err
	}
	var categoryDtoList []*Dto
	for _, category := range categoryList {
		categoryDto := &Dto{
			ID:   category.ID,
			Name: category.Name,
		}
		categoryDtoList = append(categoryDtoList, categoryDto)
	}
	return categoryDtoList, nil
}

func UpdateCategory(id int, categoryForm *Form) error {
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

func InsertCategory(categoryForm *Form) error {
	c := models.Category{
		Name: categoryForm.Name,
	}
	err := c.Insert(nil)
	if err != nil {
		return err
	}
	return nil
}
