package Models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CompanyId string `json:"id"`
	Name      string `json:"name"`
}
