package Controllers

import (
	"assignment/imdb_sql/src/Config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCompanyStat(c *gin.Context) {
	//company_id
	company_id := c.Param("id")

	whereParams := make(map[string]interface{})
	var rawStat []map[string]interface{}

	companySubQuery := Config.DB.Raw(`
		SELECT id FROM companies WHERE company_id = ?
	`, company_id)
	whereParams["company"] = companySubQuery

	movieSubQuery := Config.DB.Raw(`
		SELECT movie_id FROM movies_companies WHERE company_id in (
			SELECT id FROM companies WHERE company_id = ?
		)
	`, company_id)
	whereParams["movie"] = movieSubQuery

	//filter and get company
	ratingSubQuery := Config.DB.Raw(
		`SELECT * FROM ratings WHERE ratings.movie_id IN (?)
		`, movieSubQuery)
	whereParams["rating"] = ratingSubQuery

	if err := Config.DB.Debug().Raw(`SELECT COUNT(*) AS totalMovie, AVG(movies.imdb_rating_votes) AS imdbRatingVotes, 
		AVG(rating.imdb) AS imdbRating, AVG(rating.metacritic) AS metacritic,
		AVG(rating.the_movie_db) AS theMovieDBRating, AVG(rating.rotten_tomatoes)  AS rottenTomatoesRating,
		AVG(rating.tvcom) AS tvComRating, AVG(rating.film_affinity) AS filmAffinity
		FROM (movies, (@company) AS company, (@rating) AS rating) 
		WHERE movies.id IN (@movie) GROUP BY company.id
		`, whereParams).Find(&rawStat).Error; err == nil {
		var actorStat Stat
		for _, statRow := range rawStat {
			actorStat = Stat{
				TotalMovie:              statRow["totalMovie"].(int64),
				AvgImdbRatingVotes:      statRow["imdbRatingVotes"].(float64),
				AvgImdbRating:           statRow["imdbRating"].(float64),
				AvgMetacritic:           statRow["metacritic"].(float64),
				AvgTheMovieDBRating:     statRow["theMovieDBRating"].(float64),
				AvgRottenTomatoesRating: statRow["rottenTomatoesRating"].(float64),
				AvgTVComRating:          statRow["tvComRating"].(float64),
				AvgFilmAffinity:         statRow["filmAffinity"].(float64),
			}
		}
		c.JSON(http.StatusOK, actorStat)
	} else {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})

	}
}
