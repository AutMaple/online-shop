package models

import (
	"database/sql"

	"online.shop.autmaple.com/internal/db"
)

type Brand struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (b *Brand) QueryById(tx *sql.Tx) error {
	stmt := `select name, image from goods_brand where id = ?`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	row := prepare.QueryRow(b.ID)
	err = row.Scan(&b.Name, &b.Image)
	if err != nil {
		return DetailError(err)
	}
	return nil
}

func (b *Brand) PageQuery(tx *sql.Tx, offset, size int) ([]*Brand, error) {
	stmt := `
  SELECT id, name, image FROM goods_brand 
  WHERE id >= (SELECT id FROM goods_brand WHERE enable = true ORDER BY id LIMIT ?,1)
  AND enable = true ORDER BY id LIMIT ?`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return nil, DetailError(err)
	}
	start := (offset - 1) * size
	rows, err := prepare.Query(start, size)
	var brands []*Brand
	for rows.Next() {
		var brand Brand
		err := rows.Scan(&brand.ID, &brand.Name, &brand.Image)
		if err != nil {
			return nil, DetailError(err)
		}
		brands = append(brands, &brand)
	}
	return brands, nil
}

func (b *Brand) Insert(tx *sql.Tx) error {
	stmt := `insert into goods_brand(name, image) values(?,?)`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	result, err := prepare.Exec(b.Name, b.Image)
	if err != nil {
		return DetailError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return DetailError(err)
	}
	b.ID = int(id)
	return nil
}

func (b *Brand) Update(tx *sql.Tx) error {
	stmt := `update goods_brand set name = ?, image = ? where id = ?`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	result, err := prepare.Exec(b.Name, b.Image, b.ID)
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

func (b *Brand) Delete(tx *sql.Tx) error {
	stmt := `update goods_brand set enable = false where id = ?`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	_, err = prepare.Exec(b.ID)
	if err != nil {
		return DetailError(err)
	}
	return nil
}
