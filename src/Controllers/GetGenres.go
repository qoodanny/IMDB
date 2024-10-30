package Controllers

import (
	"assignment/imdb_sql/src/Config"
	"assignment/imdb_sql/src/Models"
	"assignment/imdb_sql/src/Utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGenres(c *gin.Context) {
	params := map[string]string{
		"orderBy":  Utils.DetermineOrderKey(c.Query("orderBy"), "genres", "id"),
		"orderKey": Utils.DetermineOrderDirection(c.Query("orderKey")),
	}
	var genres []Models.Genre
	if err := Config.DB.Debug().Raw(fmt.Sprintf(`
		SELECT DISTINCT name FROM genres ORDER BY %s %s
	`, params["orderBy"], params["orderKey"])).Find(&genres).Error; err == nil {
		c.JSON(http.StatusOK, genres)
	} else {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
	}
}
