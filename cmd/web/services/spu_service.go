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
	spu.ID = spuId
	if err != nil {
		tx.Rollback()
		return err
	}
	// 2.插入属性
	var attrIds []int
	for _, attrDto := range spuForm.Attrs {
		attr := &models.Attr{
			Attr: attrDto.Attr,
		}
		id, err := spu.InsertAttr(tx, attr)
		if err != nil {
			tx.Rollback()
			return err
		}
		attrIds = append(attrIds, id)
	}

	// 3. 建立 SPU 和 Attr 之间的关系
	err = spu.JoinAttr(tx, attrIds)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 4. 建立 Attr 与 Options 之间的关系
	for i := range attrIds {
		err := spu.AttrJoinOptions(tx, attrIds[i], spuForm.Attrs[i].Options)
		if err != nil {
			tx.Rollback()
			return err
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
	// 4. 查询属性
	attrIds, err := spu.QueryAttrId(nil)
	if err != nil {
		return nil, err
	}
	var attrDtos []*dto.AttrDto
	for _, attrId := range attrIds {
		attr := &models.Attr{ID: attrId}
		err := attr.QueryById(nil)
		if err != nil {
			return nil, err
		}
		attrDtos = append(attrDtos, &dto.AttrDto{ID: attr.ID, Attr: attr.Attr})
	}
	spuDto.Attrs = attrDtos
	// 5. 查询选项
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

// 删除应该是逻辑删除
func DeleteSpu() {

}
