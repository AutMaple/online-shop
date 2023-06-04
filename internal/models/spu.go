package models

import (
	"database/sql"
	"online.shop.autmaple.com/internal/utils/dbutil"
)

type Spu struct {
	ID         int
	Name       string
	BrandId    int
	CategoryId int
	StoreId    int
}

// QueryById may return the following error type: ErrNotRows
func (s *Spu) QueryById(tx *sql.Tx) error {
	stmt := `SELECT name, brand_id, category_id, store_id FROM goods_spu WHERE id = ? AND enable = true`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	defer prepare.Close()
	row := prepare.QueryRow(s.ID)
	err = row.Scan(&s.Name, &s.BrandId, &s.CategoryId, &s.StoreId)
	if err != nil {
		return DetailError(err)
	}
	return nil
}

func (s *Spu) PageQuery(tx *sql.Tx, offset, size int) ([]*Spu, error) {
	stmt := `SELECT id, name, brand_id, category_id FROM goods_spu
  WHERE id >= (SELECT id FROM goods_spu WHERE enable = true ORDER BY id LIMIT ?, 1) AND enable = true
  ORDER BY id LIMIT ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, DetailError(err)
	}
	defer prepare.Close()
	start := (offset - 1) * size
	rows, err := prepare.Query(start, size)
	if err != nil {
		return nil, DetailError(err)
	}
	defer rows.Close()
	var res []*Spu
	for rows.Next() {
		var spu Spu
		err := rows.Scan(&spu.ID, &spu.Name, &spu.BrandId, &spu.CategoryId)
		if err != nil {
			return nil, DetailError(err)
		}
		res = append(res, &spu)
	}
	return res, nil
}

func (s *Spu) Insert(tx *sql.Tx) (int, error) {
	stmt := `INSERT INTO goods_spu(name,brand_id,category_id, store_id) VALUES(?,?,?,?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return -1, DetailError(err)
	}
	defer prepare.Close()
	result, err := prepare.Exec(s.Name, s.BrandId, s.CategoryId, s.StoreId)
	if err != nil {
		return -1, DetailError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, DetailError(err)
	}
	return int(id), nil
}

// Delete return ErrRecordNotFound error if no rows affected
func (s *Spu) Delete(tx *sql.Tx) error {
	stmt := `UPDATE goods_spu SET enable = false WHERE id = ? and enable = true`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	result, err := prepare.Exec(s.ID)
	if err != nil {
		return DetailError(err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return DetailError(err)
	}
	if affected == 0 {
		return DetailError(ErrRecordNotFound)
	}
	return nil
}
