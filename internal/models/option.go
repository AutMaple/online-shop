package models

import (
	"database/sql"

	"online.shop.autmaple.com/internal/utils/dbutil"
)

type Option struct {
	ID     int    `json:"id"`
	AttrId int    `json:"attrId"`
	Value  string `json:"value"`
}

func (o *Option) QueryByAttrId(tx *sql.Tx) ([]*Option, error) {
	stmt := `select id, value from goods_attr_option where attr_id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, DetailError(err)
	}
	rows, err := prepare.Query(o.AttrId)
	if err != nil {
		return nil, DetailError(err)
	}
	defer rows.Close()
	var options []*Option
	for rows.Next() {
		var option Option
		err := rows.Scan(&option.ID, &option.Value)
		if err != nil {
			return nil, DetailError(err)
		}
		options = append(options, &option)
	}
	return options, nil
}

func (o *Option) Insert(tx *sql.Tx) (int, error) {
	stmt := `insert into goods_attr_option(attr_id, value) values(?, ?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return -1, DetailError(err)
	}
	result, err := prepare.Exec(o.AttrId, o.Value)
	id, err := result.LastInsertId()
	if err != nil {
		return -1, DetailError(err)
	}
	return int(id), nil
}

func (o *Option) DeleteByAttrId(tx *sql.Tx) error {
	stmt := `update goods_attr_option set enable = false where attr_id = ? and enable = true`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	_, err = prepare.Exec(o.AttrId)
	if err != nil {
		return DetailError(err)
	}
	return nil
}
