package spu

import (
	"online.shop.autmaple.com/cmd/web/service/sku"
	"online.shop.autmaple.com/internal/configs/db"
	"online.shop.autmaple.com/internal/models"
)

type Dto struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Brand    *models.Brand    `json:"brand"`
	Category *models.Category `json:"category"`
	Attrs    []*AttrDto       `json:"attrs"`
}

type Form struct {
	Name     string      `json:"name" binding:"required,min=1"`
	Brand    int         `json:"brand" binding:"required,min=1"`
	Category int         `json:"category" binding:"required,min=1"`
	Store    int         `json:"store" binding:"required,min=1"`
	Attrs    []*AttrForm `json:"attrs" binding:"required,dive,min=1"`
}

type AttrDto struct {
	ID      int          `json:"id"`
	Attr    string       `json:"attr"`
	Options []*OptionDto `json:"options"`
}

type AttrForm struct {
	Name    string   `json:"attr" binding:"required,min=1"`
	Options []string `json:"options" binding:"required,min=1"`
}

type OptionDto struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func InsertSpu(spuForm *Form) error {
	tx, err := db.GetMysqlDB().Begin()
	if err != nil {
		return err
	}
	spu := &models.Spu{
		Name:       spuForm.Name,
		BrandId:    spuForm.Brand,
		CategoryId: spuForm.Category,
		StoreId:    spuForm.Store,
	}
	// 1.插入SPU
	spuId, err := spu.Insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	spu.ID = spuId
	// 2.插入属性
	var attrIds []int
	for _, attrDto := range spuForm.Attrs {
		attr := &models.Attr{
			Name:  attrDto.Name,
			SpuID: spuId,
		}
		id, err := attr.Insert(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		attrIds = append(attrIds, id)
	}

	// 3. 建立 Attr 与 Options 之间的关系
	for i := range attrIds {
		options := spuForm.Attrs[i].Options
		for _, option := range options {
			o := &models.Option{Value: option, AttrId: attrIds[i]}
			_, err := o.Insert(tx)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

func QuerySpu(id int) (*Dto, error) {
	// 1. 查询 spu 的属性
	spuDto := &Dto{ID: id}
	spu := models.Spu{ID: id}
	err := spu.QueryById(nil)
	if err != nil {
		return nil, err
	}
	spuDto.Name = spu.Name

	// 2. 查询品牌
	brand := &models.Brand{ID: spu.BrandId}
	err = brand.QueryById(nil)
	if err != nil {
		return nil, err
	}
	spuDto.Brand = brand

	// 3. 查询分类
	category := &models.Category{ID: spu.CategoryId}
	err = category.QueryById(nil)
	if err != nil {
		return nil, err
	}
	spuDto.Category = category

	// 4. 查询属性 attr
	attr := &models.Attr{SpuID: id}
	attrList, err := attr.QueryBySpu(nil)
	if err != nil {
		return nil, err
	}
	var attrDtos []*AttrDto
	for _, attr := range attrList {
		attrDtos = append(attrDtos, &AttrDto{ID: attr.ID, Attr: attr.Name})
	}
	spuDto.Attrs = attrDtos

	// 5. 查询选项 option
	for _, attrDto := range attrDtos {
		option := &models.Option{AttrId: attrDto.ID}
		options, err := option.QueryByAttrId(nil)
		if err != nil {
			return nil, err
		}
		var optionsDto []*OptionDto
		for _, o := range options {
			dto := &OptionDto{ID: o.ID, Value: o.Value}
			optionsDto = append(optionsDto, dto)
		}
		attrDto.Options = optionsDto
	}
	return spuDto, nil
}

func PageQuerySpu(offset, size int) ([]*Dto, error) {
	spu := models.Spu{}
	spuList, err := spu.PageQuery(nil, offset, size)
	if err != nil {
		return nil, err
	}
	var spuDtoList []*Dto
	for _, spu := range spuList {
		spuDto, err := QuerySpu(spu.ID)
		if err != nil {
			return nil, err
		}
		spuDtoList = append(spuDtoList, spuDto)
	}
	return spuDtoList, err
}

func DeleteSpu(spuId int) error {
	tx, err := db.GetMysqlDB().Begin()
	if err != nil {
		return err
	}
	// 1. 删除 spu
	spu := &models.Spu{ID: spuId}
	err = spu.Delete(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 2. 删除 attr
	attr := &models.Attr{SpuID: spuId}
	// 2.1 删除前必须先获取 attrId, 用于后续删除 option
	attrIdList, err := attr.QueryIdsBySpuId(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 2.2 删除 attr
	err = attr.DeleteBySpuId(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 3. 删除 options
	option := &models.Option{}
	for _, attrId := range attrIdList {
		option.AttrId = attrId
		err := option.DeleteByAttrId(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	// 4. 删除 SKU
	skuIds, err := spu.QuerySkuIds(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, skuId := range skuIds {
		err = sku.DeleteSkuWithOuterTx(tx, skuId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
