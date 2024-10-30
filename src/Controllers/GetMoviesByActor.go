package Controllers

import (
	"assignment/imdb_sql/src/Config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMoviesByActor(c *gin.Context) {
	//star_id
	star_id := c.Param("id")

	whereParams := make(map[string]interface{})
	var rawMovies []map[string]interface{}

	movieSubQuery := Config.DB.Raw(`
		SELECT movie_id FROM movies_actors WHERE star_id in (
			SELECT id FROM stars WHERE star_id = ?
		)
	`, star_id)
	whereParams["movie"] = movieSubQuery

	actorSubQuery := Config.DB.Raw(`
		SELECT * FROM stars WHERE star_id = ?
	`, star_id)
	whereParams["actors"] = actorSubQuery

	//filter and get genre
	genreSubQuery := Config.DB.Raw(
		`SELECT * FROM genres WHERE id IN 
		(SELECT genre_id FROM movies_genres WHERE movies_genres.movie_id IN (?))
		`, movieSubQuery)
	whereParams["genres"] = genreSubQuery

	//filter and get company
	ratingSubQuery := Config.DB.Raw(
		`SELECT * FROM ratings WHERE ratings.movie_id IN (
			SELECT id FROM movies WHERE id IN (?)
		)
		`, movieSubQuery)
	whereParams["rating"] = ratingSubQuery

	//filter and get company
	boxOfficeSubQuery := Config.DB.Raw(
		`SELECT * FROM box_offices WHERE box_offices.movie_id IN (?)
		`, movieSubQuery)
	whereParams["box_office"] = boxOfficeSubQuery

	if err := Config.DB.Debug().Raw(`SELECT movies.movie_id AS id, movies.title AS title,
		movies.original_title AS originalTitle, movies.full_title AS fullTitle, movies.movie_type AS movieType,
		movies.release_year AS releaseYear, movies.image AS image, movies.release_date AS releaseDate, 
		movies.runtime_mins AS runtimeMins, movies.introduction AS introduction, 
		movies.awards AS awards, movies.imdb_rating_votes AS imdbRatingVotes,
		actors.name AS actorName, actors.image AS actorImage, 
		genres.name AS genreName, rating.year AS ratingYear, 
		rating.imdb AS imdbRating, rating.metacritic AS metacritic,
		rating.the_movie_db AS theMovieDBRating, rating.rotten_tomatoes  AS rottenTomatoesRating,
		rating.tvcom AS tvComRating, rating.film_affinity AS filmAffinity
		FROM (movies, (@actors) AS actors, (@genres) AS genres, (@rating) AS rating) WHERE movies.id IN (@movie)
		`, whereParams).Find(&rawMovies).Error; err == nil {
		var starMovies ActorMovies
		for index, movieRow := range rawMovies {
			if index == 0 {
				starMovies = ActorMovies{
					ActorName:  movieRow["actorName"].(string),
					ActorImage: movieRow["actorImage"].(string),
					Movies: []MovieRow{{
						Title:                    movieRow["title"].(string),
						OriginalTitle:            movieRow["originalTitle"].(string),
						FullTitle:                movieRow["fullTitle"].(string),
						MovieType:                movieRow["movieType"].(string),
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
						TvComRating:              movieRow["tvComRating"].(int64),
						FilmAffinity:             movieRow["filmAffinity"].(float64),
						Budget:                   "",
						OpeningWeekendUSA:        "",
						GrossUSA:                 "",
						CumulativeWorldwideGross: "",
					}},
				}
			} else {
				for index, movie := range starMovies.Movies {
					if index == len(starMovies.Movies)-1 &&
						movie.Title != movieRow["title"].(string) {
						starMovies.Movies = append(starMovies.Movies, MovieRow{
							Title:                    movieRow["title"].(string),
							OriginalTitle:            movieRow["originalTitle"].(string),
							FullTitle:                movieRow["fullTitle"].(string),
							MovieType:                movieRow["movieType"].(string),
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
							TvComRating:              movieRow["tvComRating"].(int64),
							FilmAffinity:             movieRow["filmAffinity"].(float64),
							Budget:                   "",
							OpeningWeekendUSA:        "",
							GrossUSA:                 "",
							CumulativeWorldwideGross: "",
						})
					} else if movie.Title == movieRow["title"].(string) {
						break
					}
				}
				starMovies.Movies = append(starMovies.Movies, MovieRow{
					Title:                    movieRow["title"].(string),
					OriginalTitle:            movieRow["originalTitle"].(string),
					FullTitle:                movieRow["fullTitle"].(string),
					MovieType:                movieRow["movieType"].(string),
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
					TvComRating:              movieRow["tvComRating"].(int64),
					FilmAffinity:             movieRow["filmAffinity"].(float64),
					Budget:                   "",
					OpeningWeekendUSA:        "",
					GrossUSA:                 "",
					CumulativeWorldwideGross: "",
				})
			}
		}
		c.JSON(http.StatusOK, starMovies)
	} else {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
	}
}
