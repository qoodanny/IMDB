package Controllers

import (
	"assignment/imdb_sql/src/Config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActorStat(c *gin.Context) {
	//star_id
	star_id := c.Param("id")

	whereParams := make(map[string]interface{})
	var rawStat []map[string]interface{}

	actorSubQuery := Config.DB.Raw(`
		SELECT id FROM stars WHERE star_id = ?
	`, star_id)
	whereParams["actor"] = actorSubQuery

	movieSubQuery := Config.DB.Raw(`
		SELECT movie_id FROM movies_actors WHERE star_id in (
			SELECT id FROM stars WHERE star_id = ?
		)
	`, star_id)
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
		FROM (movies, (@actor) AS actor, (@rating) AS rating) 
		WHERE movies.id IN (@movie) GROUP BY actor.id
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