package models

import (
	"database/sql"

	"online.shop.autmaple.com/internal/db"
)

type Menu struct {
	ID     int
	Name   string
	Url    string
	Icon   string
	Parent int
}

func (m *Menu) QueryById(tx *sql.Tx) error {
	stmt := `select name, url, icon, parent form ums_menu where id = ? and enable = true`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	row := prepare.QueryRow(m.ID)

	err = row.Scan(&m.Name, &m.Url, &m.Icon, &m.Parent)
	if err != nil {
		return DetailError(err)
	}
	return nil
}

func (m *Menu) PageQuery(tx *sql.Tx, offset, size int) ([]*Menu, error) {

	stmt := `
  SELECT name,url,icon,parent FROM ums_menu 
  WHERE id >=(SELECT id FROM ums_menu WHERE enable = true ORDER BY id LIMIT ?,1)
  AND enable = true
  ORDER BY id LIMIT ?`
  prepare, err := db.ToPrepare(tx, stmt)
  if err != nil {
    return nil, DetailError(err)
  }
  start := (offset - 1) * size
  rows, err := prepare.Query(start, size)
  if err != nil {
    return nil, DetailError(err)
  }
  var menus []*Menu
  for rows.Next() {
    var menu Menu
    err := rows.Scan(&menu.ID, &menu.Name, &menu.Url, &menu.Icon, &menu.Parent)
    if err != nil {
      return nil, DetailError(err)
    }
    menus = append(menus, &menu)
  }
	return menus, nil
}

func (m *Menu) Insert(tx *sql.Tx) error {
  stmt := `insert into ums_menu(name,url,icon,parent) values(?,?,?,?)`
  prepare,err := db.ToPrepare(tx, stmt)
  if err != nil {
    return DetailError(err)
  }
  result, err := prepare.Exec(&m.Name, &m.Url, &m.Icon, &m.Parent)
  if err != nil {
    return DetailError(err)
  }
  id, err := result.LastInsertId()
  if err != nil {
    return DetailError(err)
  }
  m.ID = int(id)
	return nil
}

func (m *Menu) Update(tx *sql.Tx) error {
  stmt := `update ums_menu name = ?, url = ?, icon = ?, parent = ? where id = ? and enable = true`
  prepare, err := db.ToPrepare(tx, stmt)
  if err != nil {
    return DetailError(err)
  }
  result, err := prepare.Exec(m.Name, m.Url, m.Icon, m.Parent, m.ID)
  if err != nil {
    return DetailError(err)
  }
  affected, err := result.RowsAffected()
  if err != nil {
    return DetailError(err)
  }
  if affected == 0 {
    return DetailError(sql.ErrNoRows)
  }
	return nil
}

func (m *Menu) Delete(tx *sql.Tx) error {
  stmt := `update ums_menu set enable = false where id = ? `
  prepare, err := db.ToPrepare(tx, stmt)
  if err != nil {
    return DetailError(err)
  }
  result, err := prepare.Exec(m.ID)
  if err != nil {
    return DetailError(err)
  }
  affected, err := result.RowsAffected()
  if err != nil {
    return DetailError(err)
  }
  if affected == 0 {
    return DetailError(sql.ErrNoRows)
  }
	return nil
}
