package models

import (
	"database/sql"

	"online.shop.autmaple.com/internal/utils/dbutil"
)

type Store struct {
	ID       int
	BrandIds string
	Name     string
	Address  string
	Phone    string
}

// QueryById may return the following error type: ErrNotRows
func (s *Store) QueryById(tx *sql.Tx) error {
	stmt := `select brand_ids, name, address, phone from goods_store where id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	row := prepare.QueryRow(s.ID)
  err = row.Scan(&s.BrandIds, &s.Name, &s.Address, &s.Phone);
  if err != nil {
    return DetailError(err)
  }
  return nil
}

func (s *Store) PageQuery(tx *sql.Tx, offset, size int) ([]*Store, error) {
	stmt := `SELECT id,brand_ids,name,address,phone FROM goods_store
  WHERE id >= (SELECT id FROM goods_store ORDER BY id LIMIT ?, 1) AND enable = true
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
	var res []*Store
	for rows.Next() {
		var store Store
		err := rows.Scan(&store.ID, &store.BrandIds, &store.Name, &store.Address, &store.Phone)
		if err != nil {
			return nil, DetailError(err)
		}
		res = append(res, &store)
	}
	return res, nil
}

func (s *Store) Insert(tx *sql.Tx) (int, error) {
	stmt := `INSERT INTO goods_store(brand_ids,name,address,phone) VALUES(?,?,?,?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return -1, DetailError(err)
	}
	defer prepare.Close()
	result, err := prepare.Exec(s.BrandIds, s.Name, s.Address, s.Phone)
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
func (s *Store) Delete(tx *sql.Tx) error {
	stmt := `UPDATE goods_store SET enable = false WHERE id = ? and enable = true`
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
