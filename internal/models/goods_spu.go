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
}

type Attr struct {
	ID   int
	Attr string
}

type Sku struct {
	ID    int
	SkuId int
}

func (s *Spu) QueryById(tx *sql.Tx) error {
	stmt := `select name, brand_id, category_id from goods_spu where id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return err
	}
	defer prepare.Close()
	row := prepare.QueryRow(s.ID)
	err = row.Scan(&s.Name, &s.BrandId, &s.CategoryId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Spu) PageQuery(tx *sql.Tx, offset, size int) ([]*Spu, error) {
	stmt := `select name, brand_id, category_id from goods_spu
  where id >= (select id from goods_spu order by id limit ?, 1)
  order by id limit ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, err
	}
	defer prepare.Close()
	start := (offset - 1) * size
	rows, err := prepare.Query(start, size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*Spu
	for rows.Next() {
		var spu Spu
		rows.Scan(&spu.Name, &spu.BrandId, &spu.CategoryId)
		res = append(res, &spu)
	}
	return res, nil
}

func (s *Spu) Insert(tx *sql.Tx) (int, error) {
	stmt := `insert into goods_spu(name,brand_id,category_id) 
  values(?,?,?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return -1, err
	}
	defer prepare.Close()
	result, err := prepare.Exec(s.Name, s.BrandId, s.CategoryId)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (s *Spu) InsertAttr(tx *sql.Tx, attr *Attr) (int, error) {
	stmt := `insert into goods_attr(attr) values(?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return -1, err
	}
	defer prepare.Close()
	result, err := prepare.Exec(attr.Attr)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), err
}

func (s *Spu) JoinAttr(tx *sql.Tx, attrIds []int) error {
	stmt := `insert into goods_spu_attr(spu_id, attr_id) values(?, ?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return err
	}
	defer prepare.Close()
	for _, attrId := range attrIds {
		_, err := prepare.Exec(s.ID, attrId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Spu) AttrJoinOptions(tx *sql.Tx, attrId int, options []string) error {
	stmt := `insert into goods_attr_option(attr_id, value) values(?, ?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return err
	}
	defer prepare.Close()
	for _, option := range options {
		_, err := prepare.Exec(attrId, option)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Spu) Update(tx *sql.Tx) error {
	stmt := `update goods_spu set name = ?, brand_id = ?, category_id = ? where id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.Exec(s.Name, s.BrandId, s.CategoryId, s.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Spu) Delete(tx *sql.Tx) error {
	stmt := `delete from goods_spu where id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil
	}
	_, err = prepare.Exec(s.ID)
	if err != nil {
		return err
	}
	return nil
}
