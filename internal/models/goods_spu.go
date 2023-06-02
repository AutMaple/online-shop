package models

import (
	"database/sql"
	"online.shop.autmaple.com/internal/utils/dbutil"
)

type Spu struct {
	ID         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	BrandId    int    `json:"brandId" binding:"required,min=1"`
	CategoryId int    `json:"categoryId" binding:"required,min=1"`
}

type Attr struct {
	ID    int    `json:"id"`
	SpuID int    `json:"spuId"`
	Attr  string `json:"attr" binding:"required"`
}

type Sku struct {
	ID    int
	SkuId int
}

type Brand struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Option struct {
	ID     int    `json:"id"`
	AttrId int    `json:"attrId"`
	Value  string `json:"value"`
}

func (s *Spu) QueryById(tx *sql.Tx) error {
	stmt := `SELECT name, brand_id, category_id FROM goods_spu WHERE id = ? AND enable = true`
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
	stmt := `SELECT id, name, brand_id, category_id FROM goods_spu
  WHERE id >= (SELECT id FROM goods_spu ORDER BY id LIMIT ?, 1) AND enable = true
  ORDER BY id LIMIT ?`
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
		rows.Scan(&spu.ID, &spu.Name, &spu.BrandId, &spu.CategoryId)
		res = append(res, &spu)
	}
	return res, nil
}

func (s *Spu) Insert(tx *sql.Tx) (int, error) {
	stmt := `INSERT INTO goods_spu(name,brand_id,category_id) VALUES(?,?,?)`
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
	stmt := `INSERT INTO goods_attr(attr) VALUES(?)`
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

func (a *Attr) QueryById(tx *sql.Tx) error {
	stmt := `select attr from goods_attr where id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return err
	}
	result := prepare.QueryRow(a.ID)
	err = result.Scan(&a.Attr)
	if err != nil {
		return err
	}
	return nil
}

func (s *Spu) QueryAttrId(tx *sql.Tx) ([]int, error) {
	stmt := `select attr_id from goods_spu_attr where spu_id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, err
	}
	rows, err := prepare.Query(s.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var attrIds []int
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		attrIds = append(attrIds, id)
	}
	return attrIds, nil
}

func (s *Spu) JoinAttr(tx *sql.Tx, attrIds []int) error {
	stmt := `INSERT INTO goods_spu_attr(spu_id, attr_id) VALUES(?, ?)`
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
	stmt := `INSERT INTO goods_attr_option(attr_id, value) VALUES(?, ?)`
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
	stmt := `UPDATE goods_spu SET name = ?, brand_id = ?, category_id = ? WHERE id = ? AND enable = true`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return err
	}
	defer prepare.Close()
	result, err := prepare.Exec(s.Name, s.BrandId, s.CategoryId, s.ID)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

// Delete return ErrRecordNotFound error if no rows affected
func (s *Spu) Delete(tx *sql.Tx) error {
	stmt := `UPDATE goods_spu SET enable = false WHERE id = ? and enable = true`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil
	}
	result, err := prepare.Exec(s.ID)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (b *Brand) QueryById(tx *sql.Tx) error {
	stmt := `select name, image from goods_brand where id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return err
	}
	row := prepare.QueryRow(b.ID)
	err = row.Scan(&b.Name, &b.Image)
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) QueryById(tx *sql.Tx) error {
	stmt := `select name from goods_category where id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return err
	}
	row := prepare.QueryRow(c.ID)
	err = row.Scan(&c.Name)
	if err != nil {
		return err
	}
	return nil
}

func (o *Option) QueryByAttrId(tx *sql.Tx) ([]*Option, error) {
	stmt := `select id, value from goods_attr_option where attr_id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, err
	}
	rows, err := prepare.Query(o.AttrId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var options []*Option
	for rows.Next() {
		var option Option
		err := rows.Scan(&option.ID, &option.Value)
		if err != nil {
			return nil, err
		}
		options = append(options, &option)
	}
	return options, nil
}

func (a *Attr) QueryOptionIds(tx *sql.Tx) ([]int, error) {
	stmt := `select id from goods_attr_option where attr_id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, err
	}
	rows, err := prepare.Query(a.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var optionIdList []int
	for rows.Next() {
		var optionId int
		err := rows.Scan(&optionId)
		if err != nil {
			return nil, err
		}
		optionIdList = append(optionIdList, optionId)
	}
	return optionIdList, nil
}

func (a *Attr) Delete(tx *sql.Tx) error {
	stmt := `update goods_attr set enable = false where id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(a.ID)
	if err != nil {
		return err
	}
	return nil
}

func (o *Option) Delete(tx *sql.Tx) error {
	stmt := `update goods_attr_option set enable = false where id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(o.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a *Attr) Insert(tx *sql.Tx) (int, error) {
	stmt := `insert into goods_attr(attr, spu_id) values(?,?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return -1, err
	}
	result, err := prepare.Exec(a.Attr, a.SpuID)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, nil
	}
	return int(id), nil
}

func (a *Attr) QueryBySpu(tx *sql.Tx) ([]*Attr, error) {
	stmt := `select id, attr from goods_attr where spu_id = ? and enable = true`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, err
	}
	rows, err := prepare.Query(a.SpuID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var attrList []*Attr
	for rows.Next() {
		var attr Attr
		err := rows.Scan(&attr.ID, &attr.Attr)
		if err != nil {
			return nil, err
		}
		attrList = append(attrList, &attr)
	}
	return attrList, nil
}

func (a *Attr) QueryIdsBySpuId(tx *sql.Tx) ([]int, error) {
	stmt := `select id from goods_attr where spu_id = ? and enable = true`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return nil, err
	}
	rows, err := prepare.Query(a.SpuID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var idList []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		idList = append(idList, id)
	}
	return idList, nil
}

func (a *Attr) DeleteBySpuId(tx *sql.Tx) error {
	stmt := `update goods_attr set enable = false where spu_id = ?`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(a.SpuID)
	return err
}

func (o *Option) Insert(tx *sql.Tx) (int, error) {
	stmt := `insert into goods_attr_option(attr_id, value) values(?, ?)`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return -1, err
	}
	result, err := prepare.Exec(o.AttrId, o.Value)
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (o *Option) DeleteByAttrId(tx *sql.Tx) error {
	stmt := `update goods_attr_option set enable = false where attr_id = ? and enable = true`
	prepare, err := dbutil.ToPrepare(tx, stmt)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(o.AttrId)
	return err
}
