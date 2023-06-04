package dto

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type SkuForm struct {
	Spu            int                          `json:"spu" binding:"required,min=1"`
	Stock          int                          `json:"stock" binding:"required,min=0"`
	Attrs          []int                        `json:"attrs" binding:"required,min=1"`
	Specifications map[string]map[string]string `json:"specifications" binding:"required,specifications"`
}

type SkuDto struct {
	ID             int                          `json:"id"`
	Name           string                       `json:"name"`
	Stock          int                          `json:"stock"`
	Attrs          map[string]string            `json:"attrs"`
	Specifications map[string]map[string]string `json:"specifications"`
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
