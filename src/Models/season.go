package Models

import "gorm.io/gorm"

type Season struct {
	gorm.Model
	MovieId     uint
	ImdbId      string    `json:"imDbId"`
	Title       string    `json:"title"`
	FullTitle   string    `json:"fullTitle"`
	Type        string    `json:"type"`
	ReleaseYear string    `json:"year"`
	Episodes    []Episode `json:"episodes"`
}
