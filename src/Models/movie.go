package Models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	MovieId         string    `json:"id"`
	Title           string    `json:"title"`
	OriginalTitle   string    `json:"originalTitle"`
	FullTitle       string    `json:"fullTitle"`
	MovieType       string    `json:"type"`
	ReleaseYear     int       `json:"year,string"`
	Image           string    `json:"image"`
	ReleaseDate     string    `json:"releaseDate"`
	RuntimeMins     string    `json:"runtimeMins"`
	Introduction    string    `json:"plot"`
	Awards          string    `json:"awards"`
	Actors          []Star    `json:"actorList" gorm:"many2many:movies_actors;"`
	Genres          []Genre   `json:"genreList" gorm:"many2many:movies_genres;"`
	Companies       []Company `json:"companyList" gorm:"many2many:movies_companies;"`
	Countries       []Country `json:"countryList" gorm:"many2many:movies_countries;"`
	ImdbRating      float32   `json:"imDbRating,string"`
	ImdbRatingVotes int       `json:"imDbRatingVotes,string"`
	Rating          Rating    `json:"ratings"`
	BoxOffice       BoxOffice `json:"boxOffice"`
	Seasons         []Season
}

type TVSeriesInformation struct {
	YearEnd  string   `json:"yearEnd"`
	Creators string   `json:"creators"`
	Seasons  []string `json:"seasons"`
}
type MovieExtend struct {
	Movie
	TVSeriesInfo TVSeriesInformation `json:"tvSeriesInfo"`
}
