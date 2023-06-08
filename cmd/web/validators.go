package main

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"online.shop.autmaple.com/cmd/web/service/sku"
)

func RegisterValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("specifications", sku.SpecificationValidator)
	}
}
