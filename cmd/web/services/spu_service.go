package services

import (
	"online.shop.autmaple.com/internal/configs/db"
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/models"
)

func InsertSpu(spuForm *dto.SpuForm) error {
	tx, err := db.GetMysqlDB().Begin()
	if err != nil {
		return err
	}
	spu := &models.Spu{
		Name:       spuForm.Name,
		BrandId:    spuForm.Brand,
		CategoryId: spuForm.Category,
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
			Attr:  attrDto.Attr,
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

func QuerySpu(id int) (*dto.SpuDto, error) {
	// 1. 查询 spu 的属性
	spuDto := *&dto.SpuDto{ID: id}
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
	var attrDtos []*dto.AttrDto
	for _, attr := range attrList {
		attrDtos = append(attrDtos, &dto.AttrDto{ID: attr.ID, Attr: attr.Attr})
	}
	spuDto.Attrs = attrDtos

	// 5. 查询选项 option
	for _, attrDto := range attrDtos {
		option := &models.Option{AttrId: attrDto.ID}
		options, err := option.QueryByAttrId(nil)
		if err != nil {
			return nil, err
		}
		var optionsDto []*dto.OptionDto
		for _, o := range options {
			dto := &dto.OptionDto{ID: o.ID, Value: o.Value}
			optionsDto = append(optionsDto, dto)
		}
		attrDto.Options = optionsDto
	}
	return &spuDto, nil
}

func PageQuerySpu(offset, size int) ([]*dto.SpuDto, error) {
	spu := models.Spu{}
	spuList, err := spu.PageQuery(nil, offset, size)
	if err != nil {
		return nil, err
	}
	var spuDtoList []*dto.SpuDto
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
	// TODO 4. 删除 SKU
	tx.Commit()
	return nil
}
