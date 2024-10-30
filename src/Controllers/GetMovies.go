package Controllers

import (
	"assignment/imdb_sql/src/Config"
	"assignment/imdb_sql/src/Utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	//get all query param first
	params := map[string]string{
		"movies.movie_type":        c.Query("movieType"),
		"movies.title":             c.Query("title"),
		"movies.release_year":      c.Query("releaseYear"),
		"releaseYearCompareDir":    c.Query("releaseYearCompareDir"),
		"genre":                    c.Query("genre"),
		"movies.imdb_rating":       c.Query("imdbRating"),
		"imdbRatingCompareDir":     c.Query("imdbRatingCompareDir"),
		"rating.metacritic":        c.Query("metacritic"),
		"metacriticCompareDir":     c.Query("metacriticCompareDir"),
		"rating.the_movie_db":      c.Query("theMovieDb"),
		"theMovieDbCompareDir":     c.Query("theMovieDbCompareDir"),
		"rating.rotten_tomatoes":   c.Query("rottenTomatoes"),
		"rottenTomatoesCompareDir": c.Query("rottenTomatoesCompareDir"),
		"rating.tvcom":             c.Query("tvcom"),
		"tVcomCompareDir":          c.Query("tvcomCompareDir"),
		"rating.film_affinity":     c.Query("filmAffinity"),
		"filmAffinityCompareDir":   c.Query("filmAffinityCompareDir"),
	}
	whereParams := make(map[string]interface{})
	var rawMovies []map[string]interface{}
	whereClause := " WHERE"
	for key := range params {
		if key == "movies.movie_type" {
			fmt.Println(key, " : ", params[key])
			if len(params[key]) > 0 {
				whereClause += " " + key + " = @" + key + " AND"
				whereParams[key] = params[key]
			}
		} else if key == "movies.title" {
			fmt.Println(key, " : ", params[key])
			if len(params[key]) > 0 {
				whereClause += " " + key + " LIKE @" + key + " AND"
				whereParams[key] = "%" + params[key] + "%"
			}

		} else if key == "movies.release_year" ||
			key == "movies.imdb_rating" {
			if len(params[key]) > 0 {
				switch params[Utils.DetermineSignParam(key)] {
				case "eq":
					whereClause += " " + key + " = @" + key + " AND"
					whereParams[key], _ = strconv.ParseFloat(params[key], 32)
				case "lt":
					whereClause += " " + key + " <= @" + key + " AND"
					whereParams[key], _ = strconv.ParseFloat(params[key], 32)
				default:
					whereClause += " " + key + " >= @" + key + " AND"
					whereParams[key], _ = strconv.ParseFloat(params[key], 32)
				}
			}
		}
	}
	if len(whereClause) > 6 && whereClause[len(whereClause)-3:] == "AND" {
		whereClause = whereClause[:len(whereClause)-3]
	} else if len(whereClause) == 6 {
		whereClause = " WHERE NOT movies.movie_type = @movie_type"
		whereParams["movie_type"] = ""
	}
	//filter and get genre
	genreSubQuery := Config.DB.Raw(
		fmt.Sprintf(`SELECT * FROM genres WHERE id IN 
		(SELECT genre_id FROM movies_genres WHERE movies_genres.movie_id IN (SELECT id FROM movies %s))
		`, whereClause), whereParams)
	whereParams["genres"] = genreSubQuery

	//filter and get company
	ratingSubQuery := Config.DB.Raw(
		fmt.Sprintf(`SELECT * FROM ratings WHERE ratings.movie_id IN (SELECT id FROM movies %s)
		`, whereClause), whereParams)
	whereParams["rating"] = ratingSubQuery

	//filter and get company
	boxOfficeSubQuery := Config.DB.Raw(
		fmt.Sprintf(`SELECT * FROM box_offices WHERE box_offices.movie_id IN (SELECT id FROM movies %s)
		`, whereClause), whereParams)
	whereParams["box_office"] = boxOfficeSubQuery

	if err := Config.DB.Debug().Raw(fmt.Sprintf(`SELECT movies.movie_id AS id, movies.title AS title,
		movies.original_title AS originalTitle, movies.full_title AS fullTitle, movies.movie_type AS movieType,
		movies.release_year AS releaseYear, movies.image AS image, movies.release_date AS releaseDate, 
		movies.runtime_mins AS runtimeMins, movies.introduction AS introduction, 
		movies.awards AS awards, movies.imdb_rating_votes AS imdbRatingVotes,
		genres.name AS genreName, rating.year AS ratingYear, 
		rating.imdb AS imdbRating, rating.metacritic AS metacritic,
		rating.the_movie_db AS theMovieDBRating, rating.rotten_tomatoes  AS rottenTomatoesRating,
		rating.tvcom AS tvComRating, rating.film_affinity AS filmAffinity, 
		box_office.budget AS budget, box_office.opening_weekend_usa AS openingWeekendUSA,
		box_office.gross_usa AS grossUSA, box_office.cumulative_worldwide_gross  AS cumulativeWorldwideGross
		FROM (movies, (@genres) AS genres, (@rating) AS rating, (@box_office) AS box_office) %s 
		`, whereClause), whereParams).Find(&rawMovies).Error; err == nil {
		movies := make(map[string]MovieRow)
		for _, movieRow := range rawMovies {
			if entry, ok := movies[movieRow["id"].(string)]; !ok {
				movies[movieRow["id"].(string)] = MovieRow{
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
					Budget:                   movieRow["budget"].(string),
					OpeningWeekendUSA:        movieRow["openingWeekendUSA"].(string),
					GrossUSA:                 movieRow["grossUSA"].(string),
					CumulativeWorldwideGross: movieRow["cumulativeWorldwideGross"].(string),
				}
			} else {
				entry.GenreName = append(entry.GenreName, movieRow["genreName"].(string))
			}
		}
		c.JSON(http.StatusOK, movies)
	} else {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
	}
}
