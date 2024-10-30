package Utils

func DetermineSignParam(key string) string {
	switch key {
	case "movies.release_year":
		return "releaseYearCompareDir"
	case "movies.imdb_rating":
		return "imdbRatingCompareDir"
	case "metacritic":
		return "metacriticCompareDir"
	case "theMovieDb":
		return "theMovieDbCompareDir"
	case "rottenTomatoes":
		return "rottenTomatoesCompareDir"
	case "tvcom":
		return "tvcomCompareDir"
	case "filmAffinity":
		return "filmAffinityCompareDir"
	default:
		return ""
	}
}
