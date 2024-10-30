package Models

import (
	"gorm.io/gorm"
)

type Episode struct {
	gorm.Model
	SeasonId        uint
	EpisodeId       string  `json:"id"`
	SeasonNumber    int     `json:"seasonNumber,string"`
	EpisodeNumber   int     `json:"episodeNumber,string"`
	Title           string  `json:"title"`
	Image           string  `json:"image"`
	ReleaseYear     string  `json:"year"`
	ReleaseDate     string  `json:"released"`
	Introduction    string  `json:"plot"`
	ImdbRating      float32 `json:"imDbRating,string"`
	ImdbRatingCount int     `json:"imDbRatingCount,string"`
}
