package menu

import (
	"go/format"

	"online.shop.autmaple.com/internal/models"
)

type Dto struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Icon    string `json:"icon,omitempty"`
	Url     string `json:"url,omitempty"`
	SubMenu []*Dto `json:"subMenu,omitempty"`
}

type Form struct {
	ID     int    `json:"id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Icon   string `json:"icon"`
	Url    string `json:"url"`
	Parent int    `json:"parent" bindign:"required"`
}

// TODO 通过递归的方式生成的菜单按钮，时间复杂度较高, 待优化性能
func QueryMenu() ([]*Dto, error) {
	m := models.Menu{}
	menus, err := m.QueryAll(nil)
	if err != nil {
		return nil, err
	}
	var dtos []*Dto
	for _, menu := range menus {
		if menu.Parent == -1 {
			root, err := buildMenuTree(menu, menus)
			if err != nil {
				return nil, err
			}
			dtos = append(dtos, root)
		}
	}
	return dtos, err
}

func buildMenuTree(parent *models.Menu, menus []*models.Menu) (*Dto, error) {
	pid := parent.ID
	dto := &Dto{
		ID:    pid,
		Title: parent.Name,
		Icon:  parent.Icon,
		Url:   parent.Url,
	}
	for _, menu := range menus {
		if pid == menu.Parent {
			child, err := buildMenuTree(menu, menus)
			if err != nil {
				return nil, err
			}
			dto.SubMenu = append(dto.SubMenu, child)
		}
	}
	return dto, nil
}

func InsertMenu(menu *Form) error {
	m := &models.Menu{
		Name:   menu.Title,
		Url:    menu.Url,
		Icon:   menu.Icon,
		Parent: menu.Parent,
	}
	err := m.Insert(nil)
	if err != nil {
		return err
	}
	return nil
}

func InsertMenus(menus []*Form) error {
	for _, menu := range menus {
		err := InsertMenu(menu)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateMenu(id int, menu *Form) error {
	m := &models.Menu{
		ID:     id,
		Name:   menu.Title,
		Url:    menu.Url,
		Icon:   menu.Icon,
		Parent: menu.Parent,
	}
	err := m.Update(nil)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMenu(id int) error {
	m := &models.Menu{
		ID: id,
	}
	err := m.Delete(nil)
	if err != nil {
		return err
	}
	return nil
}
