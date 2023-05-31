package services

import (
	"online.shop.autmaple.com/internal/configs/db"
	"online.shop.autmaple.com/internal/configs/log"
	"online.shop.autmaple.com/internal/models"
	"online.shop.autmaple.com/internal/dto"
)


func InsertSpu(spuDto *dto.SpuDto) {
  tx,err := db.GetMysqlDB().Begin()
  if err != nil {
    log.Error(err, "Start Transaction failed")
  }
	spu := &models.Spu{
		Name:       spuDto.Name,
		BrandId:    spuDto.Brand,
		CategoryId: spuDto.Category,
	}
	// 1.插入SPU
	spuId, err := spu.Insert()
	spu.ID = spuId
	if err != nil {
		log.Error(err, "Insert Spu Failed")
		return
	}
	// 2.插入属性
	var attrIds []int
	for _, attrDto := range spuDto.Attrs {
		attr := &models.Attr{
			Attr: attrDto.Attr,
		}
		id, err := spu.InsertAttr(attr)
		if err != nil {
			log.Error(err, "")
			return
		}
		attrIds = append(attrIds, id)
	}

	// 3. 建立 SPU 和 Attr 之间的关系
	spu.JoinAttr(attrIds)

	// 4. 建立 Attr 与 Options 之间的关系
	for i := range attrIds {
		err := spu.AttrJoinOptions(attrIds[i], spuDto.Attrs[i].Options)
		if err != nil {
			return
		}
	}
}
