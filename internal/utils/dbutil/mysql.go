package dbutil

import (
	"database/sql"

	"online.shop.autmaple.com/internal/configs/db"
)

func ToPrepare(tx *sql.Tx, stmt string) (*sql.Stmt, error) {
	var prepare *sql.Stmt
	var err error
	if tx != nil {
		prepare, err = tx.Prepare(stmt)
		if err != nil {
			return nil, err
		}
	} else {
		prepare, err = db.GetMysqlDB().Prepare(stmt)
		if err != nil {
			return nil, err
		}
	}
	return prepare, nil
}
