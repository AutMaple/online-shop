package models

import (
	"database/sql"

	"online.shop.autmaple.com/internal/configs/db"
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

func (s *Spu) QueryById(tx *sql.Tx, id int) (*Spu, error) {
	stmt := `select name, brand_id, category_id from goods_spu where id = ?`
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
	var spu *Spu
	row := prepare.QueryRow(stmt, id)
	err = row.Scan(spu.Name, spu.BrandId, spu.CategoryId)
	if err != nil {
		return nil, err
	}
	return spu, nil
}

func (s *Spu) QueryAll() {

}

func (s *Spu) PageQuery(tx *sql.Tx, offset, size int) ([]*Spu, error) {
	stmt := `select name, brand_id, category_id from goods_spu
  where id >= (select id from goods_spu order by id limit ?, 1)
  order by id limit ?`
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
	var prepare *sql.Stmt
	var err error
	if tx != nil {
		prepare, err = tx.Prepare(stmt)
		if err != nil {
			return -1, err
		}
	} else {
		prepare, err = db.GetMysqlDB().Prepare(stmt)
		if err != nil {
			return -1, err
		}
	}
	defer prepare.Close()
	result, err := prepare.Exec(stmt, s.Name, s.BrandId, s.CategoryId)
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
	var prepare *sql.Stmt
	var err error
	if tx != nil {
		prepare, err = tx.Prepare(stmt)
		if err != nil {
			return -1, err
		}
	} else {
		prepare, err = db.GetMysqlDB().Prepare(stmt)
		if err != nil {
			return -1, err
		}
	}
	defer prepare.Close()
	result, err := prepare.Exec(stmt, attr.Attr)
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
	var prepare *sql.Stmt
	var err error
	if tx != nil {
		prepare, err = tx.Prepare(stmt)
		if err != nil {
			return err
		}
	} else {
		prepare, err = db.GetMysqlDB().Prepare(stmt)
		if err != nil {
			return err
		}
	}
	defer prepare.Close()
	for _, attrId := range attrIds {
		_, err := prepare.Exec(stmt, s.ID, attrId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Spu) AttrJoinOptions(tx *sql.Tx, attrId int, options []string) error {
	stmt := `insert into goods_attr_option(attr_id, value) values(?, ?)`
	var prepare *sql.Stmt
	var err error
	if tx != nil {
		prepare, err = tx.Prepare(stmt)
		if err != nil {
			return err
		}
	} else {
		prepare, err = db.GetMysqlDB().Prepare(stmt)
		if err != nil {
			return err
		}
	}
	defer prepare.Close()
	for _, option := range options {
		_, err := prepare.Exec(stmt, attrId, option)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Spu) Update(tx *sql.Tx, id int) error {
	return nil
}

func (s *Spu) Delete(tx *sql.Tx, id int) error {
	stmt := `delete from goods_spu where id = ?`
	var prepare *sql.Stmt
	var err error
	if tx != nil {
		prepare, err = tx.Prepare(stmt)
		if err != nil {
			return err
		}
	} else {
		prepare, err = db.GetMysqlDB().Prepare(stmt)
		if err != nil {
			return err
		}
	}
	_, err = prepare.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
