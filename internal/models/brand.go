package models

import (
	"database/sql"

	"online.shop.autmaple.com/internal/utils/dbutil"
)

type Brand struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (b *Brand) QueryById(tx *sql.Tx) error {
	stmt := `select name, image from goods_brand where id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
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
