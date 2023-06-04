package models

import (
	"database/sql"

	"online.shop.autmaple.com/internal/utils/dbutil"
)

type Sku struct {
	ID    int
	SpuId int
	Stock int
}

// QueryById may return the following error thpe: ErrNotRows
func (s *Sku) QueryById(tx *sql.Tx) error {
	stmt := `SELECT spu_id, stock FROM goods_sku WHERE id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	row := prepare.QueryRow(s.ID)
	err = row.Scan(&s.SpuId, &s.Stock)
	if err != nil {
		return DetailError(err)
	}
	return nil
}

func (fs *Sku) PageQuery(tx *sql.Tx, offset, size int) ([]*Sku, error) {
	stmt := `
  SELECT id, spu_id, stock FROM goods_sku 
  WHERE id >= (SELECT id FROM goods_sku ORDER BY id LIMIT ?, 1)
  ORDER BY id
  LIMIT ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, DetailError(err)
	}
	start := (offset-1)*size + 1
	rows, err := prepare.Query(start, size)
	if err != nil {
		return nil, DetailError(err)
	}
	var res []*Sku
	for rows.Next() {
		var sku Sku
		rows.Scan(&sku.ID, &sku.SpuId, &sku.Stock)
		res = append(res, &sku)
	}
	return res, nil
}

func (s *Sku) Insert(tx *sql.Tx) error {
	stmt := `INSERT INTO goods_sku(spu_id, stock) VALUES(?,?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	result, err := prepare.Exec(s.SpuId, s.Stock)
	if err != nil {
		return DetailError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return DetailError(err)
	}
	s.ID = int(id)
	return nil
}

func (s *Sku) Update(tx *sql.Tx) error {
	stmt := `UPDATE goods_sku set sku_id = ?, stock = ? WHERE id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	_, err = prepare.Exec(s.SpuId, s.Stock, s.ID)
	if err != nil {
		return DetailError(err)
	}
	return nil
}

func (s *Sku) Delete(tx *sql.Tx) error {
	stmt := `UDPATE goods_sku SET enable = false WHERE id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	_, err = prepare.Exec(s.ID)
	if err != nil {
		return DetailError(err)
	}
	return nil
}

func (s *Sku) InsertAttrOption(tx *sql.Tx, attrs []int) error {
	stmt := `INSERT INTO goods_sku_attr_option(attr_option_id, sku_id) VALUES(?,?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	for _, option_id := range attrs {
		_, err := prepare.Exec(option_id, s.ID)
		if err != nil {
			return DetailError(err)
		}
	}
	return nil
}

func (s *Sku) InsertSpecificationGroup(tx *sql.Tx, group string) (int, error) {
	stmt := `insert into goods_sku_specification_group(sku_id, name) values(?,?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return -1, DetailError(err)
	}
	result, err := prepare.Exec(s.ID, group)
	id, err := result.LastInsertId()
	if err != nil {
		return -1, DetailError(err)
	}
	return int(id), nil
}

func (s *Sku) InsertSpecification(tx *sql.Tx, group int, name, value string) error {
	stmt := `insert into goods_sku_specification(group_id, specification, value) values(?,?,?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	_, err = prepare.Exec(group, name, value)
	if err != nil {
		return DetailError(err)
	}
	return nil
}

func (s *Sku) QueryAttrs(tx *sql.Tx) (map[string]string, error) {
	stmt := `
  SELECT a.name, b.value FROM 
  goods_sku_attr_option AS c 
  LEFT JOIN goods_attr_option AS b 
  ON c.attr_option_id = b.id 
  LEFT JOIN goods_attr AS a
  ON b.attr_id = a.id WHERE sku_id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, DetailError(err)
	}
	rows, err := prepare.Query(s.ID)
	if err != nil {
		return nil, DetailError(err)
	}
	var res = make(map[string]string)
	for rows.Next() {
		var key, value string
		err := rows.Scan(&key, &value)
		if err != nil {
			return nil, DetailError(err)
		}
		res[key] = value
	}
	return res, nil
}

func (s *Sku) QuerySpecifications(tx *sql.Tx) (map[string]map[string]string, error) {
	stmt := `SELECT a.name, b.specification, b.value FROM
  goods_sku_specification_group AS a 
  LEFT JOIN goods_sku_specification AS b
  ON a.id = b.group_id
  WHERE sku_id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, DetailError(err)
	}
	rows, err := prepare.Query(s.ID)
	if err != nil {
		return nil, DetailError(err)
	}
	var res = make(map[string]map[string]string)
	for rows.Next() {
		var group, specification, value string
		err := rows.Scan(&group, &specification, &value)
		if err != nil {
			return nil, DetailError(err)
		}
		_, ok := res[group]
		if !ok {
			res[group] = make(map[string]string)
		}
		res[group][specification] = value
	}
	return res, nil
}
