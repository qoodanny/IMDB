package Utils

import "fmt"

func DetermineOrderKey(key string, table string, defaultField string) string {
	switch key {
	case "releaseYear":
		return "movies.release_year"

	case "imdbRating":
		return "movies.imdb_rating"

	case "metacritic":
		return "rating.metacritic"

	case "theMovieDb":
		return "rating.the_movie_db"

	case "rottenTomatoes":
		return "rating.rotten_tomatoes"

	case "tvcom":
		return "rating.tvcom"

	case "filmAffinity":
		return "rating.film_affinity"

	case "genreName":
		return "genres.name"

	case "starName":
		return "stars.name"
	case "starId":
		return "stars.star_id"
	case "companyName":
		return "companies.name"
	case "companyId":
		return "companies.company_id"

	default:
		return fmt.Sprintf("%s.%s", table, defaultField)
	}
}

func DetermineOrderDirection(key string) string {
	switch key {
	case "DESC":
		return key
	default:
		return "ASC"
	}
}
