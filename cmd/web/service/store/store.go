package store

import (
	"online.shop.autmaple.com/internal/models"
)

type Dto struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type Form struct {
	BrandIds string `json:"brandIds" binding:"required,min=1"`
	Name     string `json:"name" binding:"required,min=1"`
	Address  string `json:"address" binding:"required,min=1"`
	Phone    string `json:"phone" binding:"required,min=1"`
}

func QueryStore(id int) (*Dto, error) {
	s := &models.Store{ID: id}
	err := s.QueryById(nil)
	if err != nil {
		return nil, err
	}
	store := &Dto{
		ID:      s.ID,
		Name:    s.Name,
		Address: s.Address,
		Phone:   s.Phone,
	}
	return store, nil
}

func PageQueryStore(offset, size int) ([]*Dto, error) {
	s := &models.Store{}
	storeList, err := s.PageQuery(nil, offset, size)
	if err != nil {
		return nil, err
	}
	var storeDtoList []*Dto
	for _, store := range storeList {
		storeDto := &Dto{
			ID:      store.ID,
			Name:    store.Name,
			Address: store.Address,
			Phone:   store.Phone,
		}
		storeDtoList = append(storeDtoList, storeDto)
	}
	return storeDtoList, nil
}

func InsertStore(store *Form) error {
	s := models.Store{
		BrandIds: store.BrandIds,
		Name:     store.Name,
		Address:  store.Address,
		Phone:    store.Phone,
	}
	err := s.Insert(nil)
	if err != nil {
		return err
	}
	return nil
}

func UpdateStore(id int, store *Form) error {
	s := models.Store{
		ID:       id,
		BrandIds: store.BrandIds,
		Name:     store.Name,
		Address:  store.Address,
		Phone:    store.Phone,
	}
	err := s.Update(nil)
	if err != nil {
		return err
	}
	return nil
}

func DeleteStore(id int) error {
	s := models.Store{ID: id}
	err := s.Delete(nil)
	if err != nil {
		return err
	}
	return nil
}
