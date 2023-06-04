package services

import (
	"online.shop.autmaple.com/internal/configs/db"
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/models"
)

func InsertSku(skuForm *dto.SkuForm) error {
	tx, err := db.GetMysqlDB().Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	// 1. 添加 SKU 础信息
	sku := models.Sku{SpuId: skuForm.Spu, Stock: skuForm.Stock}
	err = sku.Insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 2. 添加 sku 的属性
	err = sku.InsertAttrOption(tx, skuForm.Attrs)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 3. 添加规格组
	for group, specifications := range skuForm.Specifications {
		groupId, err := sku.InsertSpecificationGroup(tx, group)
		if err != nil {
			tx.Rollback()
			return err
		}
		// 3.1 添加规格值
		for name, value := range specifications {
			err := sku.InsertSpecification(tx, groupId, name, value)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

func QuerySku(id int) (*dto.SkuDto, error) {
	// 1. 查询 sku 的基本信息
	var skuDto dto.SkuDto
	sku := &models.Sku{ID: id}
	err := sku.QueryById(nil)
	if err != nil {
		return nil, err
	}
	// 2. 查询 sku 的名字
	spu := &models.Spu{ID: sku.SpuId}
	err = spu.QueryById(nil)
	if err != nil {
		return nil, err
	}

	// 3. 查询 sku 的属性
	attrs, err := sku.QueryAttrs(nil)
	if err != nil {
		return nil, err
	}

	// 4. 查询 sku 的规格
	specifications, err := sku.QuerySpecifications(nil)
	if err != nil {
		return nil, err
	}
	skuDto.ID = sku.ID
	skuDto.Stock = sku.Stock
	skuDto.Name = spu.Name
	skuDto.Attrs = attrs
	skuDto.Specifications = specifications
	return &skuDto, nil
}

func PageQuerySku() {

}

func UpdateSku() {

}

func DeleteSku() {

}
