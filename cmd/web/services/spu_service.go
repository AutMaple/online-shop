package services

import (
	"online.shop.autmaple.com/internal/configs/db"
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/models"
)

func InsertSpu(spuDto *dto.SpuDto) error {
	tx, err := db.GetMysqlDB().Begin()
	if err != nil {
		return err
	}
	spu := &models.Spu{
		Name:       spuDto.Name,
		BrandId:    spuDto.Brand,
		CategoryId: spuDto.Category,
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
	for _, attrDto := range spuDto.Attrs {
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
		err := spu.AttrJoinOptions(tx, attrIds[i], spuDto.Attrs[i].Options)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func QuerySpu() {
  
}

// 删除应该是逻辑删除
func DeleteSpu() {

}
