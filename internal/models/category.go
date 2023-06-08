package models

import (
	"database/sql"

	"online.shop.autmaple.com/internal/db"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Category) QueryById(tx *sql.Tx) error {
	stmt := `select name from goods_category where id = ? AND enable = true`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	row := prepare.QueryRow(c.ID)
	err = row.Scan(&c.Name)
	if err != nil {
		return DetailError(err)
	}
	return nil
}

func (c *Category) PageQuery(tx *sql.Tx, offset, size int) ([]*Category, error) {
	stmt := `
  SELECT id, name FROM goods_category 
  WHERE id >= (SELECT id FROM goods_category WHERE enable = true ORDER BY id LIMIT ?, 1) AND enable = true
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
	var categorys []*Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, DetailError(err)
		}
		categorys = append(categorys, &category)
	}
	return categorys, nil
}

func (c *Category) Insert(tx *sql.Tx) error {
	stmt := `INSERT INTO goods_category(name) VALUES(?)`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	result, err := prepare.Exec(c.Name)
	if err != nil {
		return DetailError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return DetailError(err)
	}
	c.ID = int(id)
	return nil
}

func (c *Category) Update(tx *sql.Tx) error {
	stmt := `UPDATE goods_category SET name = ? WHERE id = ?`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	result, err := prepare.Exec(c.Name, c.ID)
	if err != nil {
		return DetailError(err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return DetailError(err)
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (c *Category) Delete(tx *sql.Tx) error {
	stmt := `UPDATE goods_category SET enable = false WHERE id = ?`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	result, err := prepare.Exec(c.ID)
	if err != nil {
		return DetailError(err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return DetailError(err)
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
