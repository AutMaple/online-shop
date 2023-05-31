package models

import (
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

func (s *Spu) QueryById(id int) (*Spu, error) {
	stmt := `select name, brand_id, category_id from goods_spu where id = ?`
	var spu *Spu
	row := db.GetMysqlDB().QueryRow(stmt, id)
	err := row.Scan(spu.Name, spu.BrandId, spu.CategoryId)
	if err != nil {
		return nil, err
	}
	return spu, nil
}

func (s *Spu) QueryAll() {

}

func (s *Spu) PageQuery() {

}

func (s *Spu) Insert() (int, error) {
	stmt := `insert into goods_spu(name,brand_id,category_id) 
  values(?,?,?)`
  prepare, err := db.GetMysqlDB().Prepare(stmt)
  if err != nil {
    return -1, err
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

func (s *Spu) InsertAttr(attr *Attr) (int, error) {
	stmt := `insert into goods_attr(attr) values(?)`
  prepare,err := db.GetMysqlDB().Prepare(stmt)
  if err != nil{
    return -1, err
  }
  prepare.Close()
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

func (s *Spu) JoinAttr(attrIds []int) error {
	stmt := `insert into goods_spu_attr(spu_id, attr_id) values(?, ?)`
  prepare, err := db.GetMysqlDB().Prepare(stmt)
  if err != nil{
    return err
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

func (s *Spu) AttrJoinOptions(attrId int, options []string) error {
	stmt := `insert into goods_attr_option(attr_id, value) values(?, ?)`
  prepare, err := db.GetMysqlDB().Prepare(stmt)
  if err != nil {
    return err
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

func (s *Spu) Update() {

}

func (s *Spu) Delete() {

}
