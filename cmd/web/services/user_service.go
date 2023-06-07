package services

import (
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/models"
)

func QueryUser(id int) (*dto.UserDto, error) {
	u := models.User{ID: id}
	err := u.QueryById(nil)
	if err != nil {
		return nil, err
	}
	user := &dto.UserDto{
		Name:  u.Name,
		Email: u.Email,
		Phone: u.Phone,
	}
	return user, nil
}

func PageQueryUser(offset, size int) ([]*dto.UserDto, error) {
	u := &models.User{}
	users, err := u.PageQuery(nil, offset, size)
	if err != nil {
		return nil, err
	}
	var userDtos []*dto.UserDto
	for _, user := range users {
		dto := &dto.UserDto{
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		}
		userDtos = append(userDtos, dto)
	}
	return userDtos, nil
}

func InsertUser(user *dto.UserForm) error {
	u := models.User{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}
	err := u.Insert(nil)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(id int, user *dto.UserForm) error {
	u := models.User{
		ID:    id,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}
	err := u.Update(nil)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	u := models.User{
		ID: id,
	}
	err := u.Delete(nil)
	if err != nil {
		return err
	}
	return nil
}
