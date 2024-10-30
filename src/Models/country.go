package Models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name string `json:"value"`
}
