package Controllers

type MovieRow struct {
	Title                    string   `json:"title"`
	OriginalTitle            string   `json:"originalTitle"`
	FullTitle                string   `json:"fullTitle"`
	MovieType                string   `json:"movieType"`
	ReleaseYear              int64    `json:"releaseYear"`
	Image                    string   `json:"image"`
	ReleaseDate              string   `json:"releaseDate"`
	RuntimeMins              string   `json:"runtimeMins"`
	Introduction             string   `json:"introduction"`
	Awards                   string   `json:"awards"`
	ImdbRatingVotes          int64    `json:"imdbRatingVotes: number,"`
	GenreName                []string `json:"genres"`
	RatingYear               int64    `json:"ratingYear"`
	ImdbRating               float64  `json:"imdbRating"`
	Metacritic               int64    `json:"metacritic"`
	TheMovieDBRating         float64  `json:"theMovieDBRating"`
	RottenTomatoesRating     int64    `json:"rottenTomatoesRating"`
	TvComRating              int64    `json:"tvComRating"`
	FilmAffinity             float64  `json:"filmAffinity"`
	Budget                   string   `json:"budget"`
	OpeningWeekendUSA        string   `json:"openingWeekendUSA"`
	GrossUSA                 string   `json:"grossUSA"`
	CumulativeWorldwideGross string   `json:"cumulativeWorldwideGross"`
}

type ActorMovies struct {
	Movies     []MovieRow `json:"movies"`
	ActorName  string     `json:"actorName"`
	ActorImage string     `json:"actorImage"`
}

type CompanyMovies struct {
	Movies      []MovieRow `json:"movies"`
	CompanyName string     `json:"companyName"`
	CompanyId   string     `json:"companyId"`
}

type ActorRow struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
type MovieDetail struct {
	Title                    string     `json:"title"`
	OriginalTitle            string     `json:"originalTitle"`
	FullTitle                string     `json:"fullTitle"`
	MovieType                string     `json:"movieType"`
	ReleaseYear              int64      `json:"releaseYear"`
	Image                    string     `json:"image"`
	ReleaseDate              string     `json:"releaseDate"`
	RuntimeMins              string     `json:"runtimeMins"`
	Introduction             string     `json:"introduction"`
	Awards                   string     `json:"awards"`
	ImdbRatingVotes          int64      `json:"imdbRatingVotes"`
	GenreName                []string   `json:"genres"`
	RatingYear               int64      `json:"ratingYear"`
	ImdbRating               float64    `json:"imdbRating"`
	Metacritic               int64      `json:"metacritic"`
	TheMovieDBRating         float64    `json:"theMovieDBRating"`
	RottenTomatoesRating     int64      `json:"rottenTomatoesRating"`
	TvComRating              float64    `json:"tvComRating"`
	FilmAffinity             float64    `json:"filmAffinity"`
	Budget                   string     `json:"budget"`
	OpeningWeekendUSA        string     `json:"openingWeekendUSA"`
	GrossUSA                 string     `json:"grossUSA"`
	CumulativeWorldwideGross string     `json:"cumulativeWorldwideGross"`
	Actors                   []ActorRow `json:"actors"`
	Companies                []string   `json:"companies"`
	Countries                []string   `json:"countries"`
}
type Stat struct {
	TotalMovie              int64   `json:"totalMovie"`
	AvgImdbRatingVotes      float64 `json:"avgImdbRatingVotes"`
	AvgImdbRating           float64 `json:"avgImdbRating"`
	AvgMetacritic           float64 `json:"avgMetacritic"`
	AvgTheMovieDBRating     float64 `json:"avgTheMovieDBRating"`
	AvgRottenTomatoesRating float64 `json:"avgRottenTomatoesRating"`
	AvgTVComRating          float64 `json:"avgTVComRating"`
	AvgFilmAffinity         float64 `json:"avgFilmAffinity"`
}
