package models

import (
	"database/sql"

	"online.shop.autmaple.com/internal/utils/dbutil"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Category) QueryById(tx *sql.Tx) error {
	stmt := `select name from goods_category where id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
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
