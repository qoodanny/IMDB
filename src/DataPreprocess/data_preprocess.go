package DataPreprocess

import (
	"assignment/imdb_sql/src/Config"
	"assignment/imdb_sql/src/Constants"
	"assignment/imdb_sql/src/Models"
	"fmt"
	"sync"
)

type Movie struct {
	Id              string `json:"id"`
	Rank            string `json:"rank"`
	Title           string `json:"title"`
	FullTitle       string `json:"fullTitle"`
	Year            string `json:"releaseYear"`
	Crew            string `json:"crew"`
	ImdbRating      string `json:"imDbRating"`
	ImdbRatingCOunt string `json:"imDbRatingCount"`
}

type GetTopMoviesResp struct {
	Items  []Movie `json:"items"`
	ErrMsg string  `json:"errorMessage"`
}

func GetSeason(id string, seasonNo string, wg *sync.WaitGroup, season *Models.Season) {
	defer wg.Done()
	if err := DataFetcher(Constants.BASE_URL+"/SeasonEpisodes/"+Constants.API_KEY+"/"+id+"/"+seasonNo, &season); err == nil {
		Config.DB.Create(&season)
	}
}
func GetMovie(id string, wg *sync.WaitGroup) {
	defer wg.Done()
	var movieExtend Models.MovieExtend
	fmt.Println("GetMovie", id)
	if err := DataFetcher(Constants.BASE_URL+"/Title/"+Constants.API_KEY+"/"+id+"/FullActor,FullCast,Ratings", &movieExtend); err == nil {
		fmt.Println("GetMovie", movieExtend.FullTitle)
		var movie = Models.Movie{
			MovieId:         movieExtend.MovieId,
			Title:           movieExtend.Title,
			OriginalTitle:   movieExtend.OriginalTitle,
			FullTitle:       movieExtend.FullTitle,
			MovieType:       movieExtend.MovieType,
			ReleaseYear:     movieExtend.ReleaseYear,
			Image:           movieExtend.Image,
			ReleaseDate:     movieExtend.ReleaseDate,
			RuntimeMins:     movieExtend.RuntimeMins,
			Introduction:    movieExtend.Introduction,
			Awards:          movieExtend.Awards,
			Actors:          movieExtend.Actors,
			Genres:          movieExtend.Genres,
			Companies:       movieExtend.Companies,
			Countries:       movieExtend.Countries,
			ImdbRating:      movieExtend.ImdbRating,
			ImdbRatingVotes: movieExtend.ImdbRatingVotes,
			Rating:          movieExtend.Rating,
			BoxOffice:       movieExtend.BoxOffice,
			Seasons:         movieExtend.Seasons,
		}
		if movieExtend.TVSeriesInfo.Seasons != nil {
			var seasonWG sync.WaitGroup
			for _, seasonData := range movieExtend.TVSeriesInfo.Seasons {
				var season Models.Season
				seasonWG.Add(1)
				go GetSeason(id, seasonData, &seasonWG, &season)
				movie.Seasons = append(movie.Seasons, season)
			}
			seasonWG.Wait()
		}
		Config.DB.Create(&movie)
	} else {
		fmt.Println(err)
	}
}

func Get250TopMovies() {
	var movies GetTopMoviesResp
	var wg sync.WaitGroup
	if err := DataFetcher(Constants.BASE_URL+"/Top250TVs/"+Constants.API_KEY, &movies); err == nil {
		for index, movie := range movies.Items {
			//we need to limit number of movie to be loaded due to free APIKEY limitation
			if index < 15 {
				fmt.Println("index", index, movie.Id)
				wg.Add(1)
				go GetMovie(movie.Id, &wg)
				wg.Wait()
			}
		}
	} else {
		fmt.Println(err)
	}
	if err := DataFetcher(Constants.BASE_URL+"/Top250Movies/"+Constants.API_KEY, &movies); err == nil {
		for index, movie := range movies.Items {
			//we need to limit number of movie to be loaded due to free APIKEY limitation
			if index < 12 {
				fmt.Println("index", index, movie.Id)
				wg.Add(1)
				go GetMovie(movie.Id, &wg)
				wg.Wait()
			}
		}
	} else {
		fmt.Println(err)
	}
}
