package main

import (
	"assignment/imdb_sql/src/Config"
	"assignment/imdb_sql/src/DataPreprocess"
	"assignment/imdb_sql/src/Models"
	"assignment/imdb_sql/src/Routes"

	"os"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

func main() {
	var err error
	if Config.DB, err = gorm.Open(sqlite.Open("movie_data.db"), &gorm.Config{}); err == nil {
		Config.DB.AutoMigrate(&Models.Movie{}, &Models.Rating{}, &Models.BoxOffice{}, &Models.Season{}, &Models.Episode{})
		if len(os.Args) > 1 {
			DataPreprocess.Get250TopMovies()
		} else {
			r := Routes.SetupRouter()
			// Listen and Server in 0.0.0.0:8080
			r.Run(":8080")
		}
	}

}
