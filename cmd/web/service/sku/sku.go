package sku

import (
	"database/sql"
	"strings"

	"github.com/go-playground/validator/v10"
	"online.shop.autmaple.com/internal/configs/db"
	"online.shop.autmaple.com/internal/models"
)
type Form struct {
	Spu            int                          `json:"spu" binding:"required,min=1"`
	Stock          int                          `json:"stock" binding:"required,min=0"`
	Attrs          []int                        `json:"attrs" binding:"required,min=1"`
	Specifications map[string]map[string]string `json:"specifications" binding:"required,specifications"`
}

type Dto struct {
	ID             int                          `json:"id"`
	Name           string                       `json:"name"`
	Stock          int                          `json:"stock"`
	Attrs          map[string]string            `json:"attrs"`
	Specifications map[string]map[string]string `json:"specifications"`
}
func InsertSku(skuForm *Form) error {
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

func QuerySku(id int) (*Dto, error) {
	// 1. 查询 sku 的基本信息
	var skuDto Dto
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

func PageQuerySku(offset, size int) ([]*Dto, error) {
	sku := models.Sku{}
	skus, err := sku.PageQuery(nil, offset, size)
	if err != nil {
		return nil, err
	}
	var res []*Dto
	for _, sku := range skus {
		err := sku.QueryById(nil)
		if err != nil {
			return nil, err
		}
		skuDto, err := QuerySku(sku.ID)
		if err != nil {
			return nil, err
		}
		res = append(res, skuDto)
	}
	return res, nil
}


func UpdateSku(id int, skuForm *Form) error {
	tx, err := db.GetMysqlDB().Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	sku := &models.Sku{ID: id}
	err = sku.QueryById(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 1. 删除
	// 1.1 删除属性 attr
	err = sku.DeleteAttrs(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 1.2 删除规格值
	err = sku.DeleteSpecification(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 2. 更新
	// 2.1 插入属性
	err = sku.InsertAttrOption(tx, skuForm.Attrs)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 2.2 插入规格值
	for group, specifications := range skuForm.Specifications {
		groupId, err := sku.InsertSpecificationGroup(tx, group)
		if err != nil {
			tx.Rollback()
			return err
		}
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

func DeleteSku(id int) error {
	tx, err := db.GetMysqlDB().Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	sku := &models.Sku{ID: id}
	err = sku.QueryById(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 1. 删除 SKU
	err = sku.Delete(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 2. 删除属性表
	err = sku.DeleteAttrs(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 3. 删除规格表
	err = sku.DeleteSpecification(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DeleteSkuWithOuterTx(tx *sql.Tx, id int) error {
	sku := &models.Sku{ID: id}
	err := sku.QueryById(tx)
	if err != nil {
		return err
	}
	// 1. 删除 SKU
	err = sku.Delete(tx)
	if err != nil {
		return err
	}
	// 2. 删除属性表
	err = sku.DeleteAttrs(tx)
	if err != nil {
		return err
	}
	// 3. 删除规格表
	err = sku.DeleteSpecification(tx)
	if err != nil {
		return err
	}
	return nil
}


// SpecificationValidator require the key and the value of specification not empty
func SpecificationValidator(fl validator.FieldLevel) bool {
	specificationGroups, ok := fl.Field().Interface().(map[string]map[string]string)
	if ok {
		if groupLen := len(specificationGroups); groupLen == 0 {
			return false
		}
		for _, specifications := range specificationGroups {
			if specLen := len(specifications); specLen == 0 {
				return false
			}
			for name, value := range specifications {
				if len(strings.TrimSpace(name)) == 0 || len(strings.TrimSpace(value)) == 0 {
					return false
				}
			}
		}
		return true
	}
	return true
}
