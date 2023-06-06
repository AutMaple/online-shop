package services

import (
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/models"
)

func QueryStore(id int) (*dto.StoreDto, error) {
	s := &models.Store{ID: id}
	err := s.QueryById(nil)
	if err != nil {
		return nil, err
	}
	store := &dto.StoreDto{
		ID:      s.ID,
		Name:    s.Name,
		Address: s.Address,
		Phone:   s.Phone,
	}
	return store, nil
}

func PageQueryStore(offset, size int) ([]*dto.StoreDto, error) {
	s := &models.Store{}
	storeList, err := s.PageQuery(nil, offset, size)
	if err != nil {
		return nil, err
	}
	var storeDtoList []*dto.StoreDto
	for _, store := range storeList {
		storeDto := &dto.StoreDto{
			ID:      store.ID,
			Name:    store.Name,
			Address: store.Address,
			Phone:   store.Phone,
		}
		storeDtoList = append(storeDtoList, storeDto)
	}
	return storeDtoList, nil
}

func InsertStore(store *dto.StoreForm) error {
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

func UpdateStore(id int, store *dto.StoreForm) error {
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
