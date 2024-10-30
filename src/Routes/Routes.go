package Routes

import (
	"assignment/imdb_sql/src/Controllers"

	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func createHTMLRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", "templates/layouts/base.html", "templates/home/index.html")
	return r
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())
	r.Use(static.Serve("/static", static.LocalFile("./static", false)))
	r.HTMLRender = createHTMLRender()
	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//home page
	r.GET("/", Controllers.HomeIndex)

	//get movie list
	r.GET("/movies", Controllers.GetMovies)
	//get movie detail
	r.GET("/movies/:id", Controllers.GetMovieDetail)
	//get actors
	r.GET("/actors", Controllers.GetActors)
	//get movie by actor
	r.GET("/actors/movies/:id", Controllers.GetMoviesByActor)
	//get actor stat
	r.GET("/actors/:id", Controllers.GetActorStat)
	//get companies
	r.GET("/companies", Controllers.GetCompanies)
	//get movie by company
	r.GET("/companies/movies/:id", Controllers.GetMoviesByCompany)
	//get company stat
	r.GET("/companies/:id", Controllers.GetCompanyStat)
	//get genres
	r.GET("/genres", Controllers.GetGenres)
	//get genre stat
	r.GET("/genres/:name", Controllers.GetGenreStat)
	return r
}
