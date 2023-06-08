package user

import (
	"online.shop.autmaple.com/internal/models"
)

type Dto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Form struct {
	Name  string `json:"name" binding:"required,min=1"`
	Email string `json:"email" binding:"required,min=1"`
	Phone string `json:"phone" binding:"required,min=1"`
}

func QueryUser(id int) (*Dto, error) {
	u := models.User{ID: id}
	err := u.QueryById(nil)
	if err != nil {
		return nil, err
	}
	user := &Dto{
		Name:  u.Name,
		Email: u.Email,
		Phone: u.Phone,
	}
	return user, nil
}

func PageQueryUser(offset, size int) ([]*Dto, error) {
	u := &models.User{}
	users, err := u.PageQuery(nil, offset, size)
	if err != nil {
		return nil, err
	}
	var userDtos []*Dto
	for _, user := range users {
		dto := &Dto{
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		}
		userDtos = append(userDtos, dto)
	}
	return userDtos, nil
}

func InsertUser(user *Form) error {
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

func UpdateUser(id int, user *Form) error {
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
