package models

import (
	"database/sql"

	"online.shop.autmaple.com/internal/db"
)

type User struct {
	ID    int
	Name  string
	Email string
	Phone string
}

func (u *User) QueryById(tx *sql.Tx) error {
	stmt := `SELECT name, email, phone FROM ums_user WHERE id = ? and enable = true`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	row := prepare.QueryRow(u.ID)
	err = row.Scan(&u.Name, &u.Email, &u.Phone)
	if err != nil {
		return DetailError(err)
	}
	return nil
}

func (u *User) PageQuery(tx *sql.Tx, offset, size int) ([]*User, error) {
	stmt := `
  SELECT id, name, email, phone FROM ums_user 
  WHERE id >= (SELECT id FROM ums_user WHERE enable = true ORDER BY id LIMIT ?,1) 
  AND enable = true ORDER BY id LIMIT ?`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return nil, DetailError(err)
	}
	start := (offset - 1) * size
	rows, err := prepare.Query(start, size)
	if err != nil {
		return nil, err
	}
	var users []*User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone)
		if err != nil {
			return nil, DetailError(err)
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *User) Insert(tx *sql.Tx) error {
	stmt := `insert into ums_user(name, email, phone) values(?,?,?)`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	result, err := prepare.Exec(u.Name, u.Email, u.Phone)
	if err != nil {
		return DetailError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return DetailError(err)
	}
	u.ID = int(id)
	return nil
}

func (u *User) Update(tx *sql.Tx) error {
	stmt := `update ums_user set name = ?, email = ?, phone = ? where id = ?`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	result, err := prepare.Exec(u.Name, u.Email, u.Phone, u.ID)
	if err != nil {
		return DetailError(err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return DetailError(err)
	}
	if affected == 0 {
		return DetailError(sql.ErrNoRows)
	}
	return nil
}

func (u *User) Delete(tx *sql.Tx) error {
	stmt := `update ums_user set enable = false where id = ?`
	prepare, err := db.ToPrepare(tx, stmt)
	if err != nil {
		return DetailError(err)
	}
	result, err := prepare.Exec(u.ID)
	if err != nil {
		return DetailError(err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return DetailError(err)
	}
	if affected == 0 {
		return DetailError(sql.ErrNoRows)
	}
	return nil
}
