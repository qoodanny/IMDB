package Models

import "gorm.io/gorm"

type BoxOffice struct {
	gorm.Model
	MovieId                  uint
	Budget                   string `json:"budget"`
	OpeningWeekendUSA        string `json:"openingWeekendUSA"`
	GrossUSA                 string `json:"grossUSA"`
	CumulativeWorldwideGross string `json:"cumulativeWorldwideGross"`
}
