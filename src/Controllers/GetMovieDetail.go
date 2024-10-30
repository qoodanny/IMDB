package Controllers

import (
	"assignment/imdb_sql/src/Config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovieDetail(c *gin.Context) {
	//movie_id
	id := c.Param("id")
	whereParams := make(map[string]interface{})
	var rawMovies []map[string]interface{}
	whereParams["id"] = id

	//filter and get genre
	actorsSubQuery := Config.DB.Raw(
		`SELECT * FROM stars WHERE id IN 
		(SELECT star_id FROM movies_actors WHERE movies_actors.movie_id = (
			SELECT id FROM movies WHERE movie_id = ?
		))
		`, id)
	whereParams["actors"] = actorsSubQuery

	//filter and get genre
	genreSubQuery := Config.DB.Raw(
		`SELECT * FROM genres WHERE id IN 
		(SELECT genre_id FROM movies_genres WHERE movies_genres.movie_id = (
			SELECT id FROM movies WHERE movie_id = ?
		))
		`, id)
	whereParams["genres"] = genreSubQuery

	//filter and get company
	companiesSubQuery := Config.DB.Raw(
		`SELECT * FROM companies WHERE id IN 
		(SELECT company_id FROM movies_companies WHERE movies_companies.movie_id = (
			SELECT id FROM movies WHERE movie_id = ?
		))
		`, id)
	whereParams["companies"] = companiesSubQuery

	//filter and get genre
	countriesSubQuery := Config.DB.Raw(
		`SELECT * FROM countries WHERE id IN 
		(SELECT country_id FROM movies_countries WHERE movies_countries.movie_id = (
			SELECT id FROM movies WHERE movie_id = ?
		))
		`, id)
	whereParams["countries"] = countriesSubQuery

	//filter and get company
	ratingSubQuery := Config.DB.Raw(
		`SELECT * FROM ratings WHERE ratings.movie_id = (
			SELECT id FROM movies WHERE movie_id = ?
		)
		`, id)
	whereParams["rating"] = ratingSubQuery

	//filter and get company
	boxOfficeSubQuery := Config.DB.Raw(
		`SELECT * FROM box_offices WHERE box_offices.movie_id = (
			SELECT id FROM movies WHERE movie_id = ?
		)
		`, id)
	whereParams["box_office"] = boxOfficeSubQuery

	if err := Config.DB.Debug().Raw(`SELECT movies.movie_id AS id, movies.title AS title,
		movies.original_title AS originalTitle, movies.full_title AS fullTitle, movies.movie_type AS movieType,
		movies.release_year AS releaseYear, movies.image AS image, movies.release_date AS releaseDate, 
		movies.runtime_mins AS runtimeMins, movies.introduction AS introduction, 
		movies.awards AS awards, movies.imdb_rating_votes AS imdbRatingVotes,
		actors.name AS actorName, actors.image AS actorImage, genres.name AS genreName,
		companies.name AS companyName, countries.name AS countryName, 
		rating.year AS ratingYear, rating.imdb AS imdbRating, rating.metacritic AS metacritic,
		rating.the_movie_db AS theMovieDBRating, rating.rotten_tomatoes  AS rottenTomatoesRating,
		rating.tvcom AS tvComRating, rating.film_affinity AS filmAffinity,
		box_office.budget AS budget, box_office.opening_weekend_usa AS openingWeekendUSA, 
		box_office.gross_usa AS grossUSA, box_office.cumulative_worldwide_gross AS cumulativeWorldwideGross
		FROM (movies, (@actors) AS actors, (@genres) AS genres, 
		(@companies) AS companies, (@countries) AS countries, 
		(@box_office) AS box_office,(@rating) AS rating) WHERE movies.movie_id = @id
		`, whereParams).Find(&rawMovies).Error; err == nil {

		var movieDetail MovieDetail
		for index, movieRow := range rawMovies {
			if index == 0 {
				movieDetail = MovieDetail{
					Title:                    movieRow["title"].(string),
					OriginalTitle:            movieRow["originalTitle"].(string),
					FullTitle:                movieRow["fullTitle"].(string),
					ReleaseYear:              movieRow["releaseYear"].(int64),
					Image:                    movieRow["image"].(string),
					ReleaseDate:              movieRow["releaseDate"].(string),
					RuntimeMins:              movieRow["runtimeMins"].(string),
					Introduction:             movieRow["introduction"].(string),
					Awards:                   movieRow["awards"].(string),
					ImdbRatingVotes:          movieRow["imdbRatingVotes"].(int64),
					GenreName:                []string{movieRow["genreName"].(string)},
					RatingYear:               movieRow["ratingYear"].(int64),
					ImdbRating:               movieRow["imdbRating"].(float64),
					Metacritic:               movieRow["metacritic"].(int64),
					TheMovieDBRating:         movieRow["theMovieDBRating"].(float64),
					RottenTomatoesRating:     movieRow["rottenTomatoesRating"].(int64),
					Budget:                   movieRow["budget"].(string),
					OpeningWeekendUSA:        movieRow["openingWeekendUSA"].(string),
					GrossUSA:                 movieRow["grossUSA"].(string),
					CumulativeWorldwideGross: movieRow["cumulativeWorldwideGross"].(string),
					Actors: []ActorRow{{
						Name:  movieRow["actorName"].(string),
						Image: movieRow["actorImage"].(string),
					}},
					Companies: []string{movieRow["companyName"].(string)},
					Countries: []string{movieRow["countryName"].(string)},
				}
			} else {
				for index, actor := range movieDetail.Actors {
					if index == len(movieDetail.Actors)-1 &&
						actor.Name != movieRow["actorName"].(string) {
						movieDetail.Actors = append(movieDetail.Actors, ActorRow{
							Name:  movieRow["actorName"].(string),
							Image: movieRow["actorImage"].(string),
						})
					} else if actor.Name == movieRow["actorName"].(string) {
						break
					}
				}
				for index, company := range movieDetail.Companies {
					if index == len(movieDetail.Companies)-1 &&
						company != movieRow["companyName"].(string) {
						movieDetail.Companies = append(movieDetail.Companies, movieRow["companyName"].(string))
					} else if company == movieRow["companyName"].(string) {
						break
					}
				}
				for index, country := range movieDetail.Countries {
					if index == len(movieDetail.Companies)-1 &&
						country != movieRow["countryName"].(string) {
						movieDetail.Companies = append(movieDetail.Countries, movieRow["countryName"].(string))
					} else if country == movieRow["countryName"].(string) {
						break
					}
				}
			}

		}
		c.JSON(http.StatusOK, movieDetail)
	} else {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
	}
}
