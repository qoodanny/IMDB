package Models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	MovieId        uint
	ImdbId         string  `json:"imDbId"`
	Title          string  `json:"title"`
	FullTitle      string  `json:"fullTitle"`
	Type           string  `json:"type"`
	Year           int     `json:"year,string"`
	Imdb           float32 `json:"imDb,string" `
	Metacritic     int     `json:"metacritic,string"`
	TheMovieDB     float32 `json:"theMovieDb,string"`
	RottenTomatoes int     `json:"rottenTomatoes,string"`
	TVCOM          float32 `json:"tV_com,string" gorm:"type:decimal(5,2)"`
	FilmAffinity   float32 `json:"filmAffinity,string" gorm:"type:decimal(5,2)"`
}
