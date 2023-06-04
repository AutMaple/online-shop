package models

import (
	"database/sql"

	"online.shop.autmaple.com/internal/utils/dbutil"
)

type Attr struct {
	ID    int    `json:"id"`
	SpuID int    `json:"spuId"`
	Name  string `json:"attr" binding:"required"`
}

func (a *Attr) Insert(tx *sql.Tx) (int, error) {
	stmt := `insert into goods_attr(name, spu_id) values(?,?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return -1, DetailError(err)
	}
	result, err := prepare.Exec(a.Name, a.SpuID)
	if err != nil {
		return -1, DetailError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, DetailError(err)
	}
	return int(id), nil
}

func (a *Attr) QueryBySpu(tx *sql.Tx) ([]*Attr, error) {
	stmt := `select id, name from goods_attr where spu_id = ? and enable = true`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, DetailError(err)
	}
	rows, err := prepare.Query(a.SpuID)
	if err != nil {
		return nil, DetailError(err)
	}
	defer rows.Close()
	var attrList []*Attr
	for rows.Next() {
		var attr Attr
		err := rows.Scan(&attr.ID, &attr.Name)
		if err != nil {
			return nil, DetailError(err)
		}
		attrList = append(attrList, &attr)
	}
	return attrList, nil
}

func (a *Attr) QueryIdsBySpuId(tx *sql.Tx) ([]int, error) {
	stmt := `select id from goods_attr where spu_id = ? and enable = true`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, DetailError(err)
	}
	rows, err := prepare.Query(a.SpuID)
	if err != nil {
		return nil, DetailError(err)
	}
	defer rows.Close()
	var idList []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, DetailError(err)
		}
		idList = append(idList, id)
	}
	return idList, nil
}

func (a *Attr) DeleteBySpuId(tx *sql.Tx) error {
	stmt := `update goods_attr set enable = false where spu_id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	_, err = prepare.Exec(a.SpuID)
	if err != nil {
		return DetailError(err)
	}
	return nil
}
