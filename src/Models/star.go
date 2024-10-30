package Models

import "gorm.io/gorm"

type Star struct {
	gorm.Model
	StarId      string `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	AsCharacter string `json:"asCharacter"`
}
