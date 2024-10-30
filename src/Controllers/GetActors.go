package Controllers

import (
	"assignment/imdb_sql/src/Config"
	"assignment/imdb_sql/src/Models"
	"assignment/imdb_sql/src/Utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActors(c *gin.Context) {
	params := map[string]string{
		"orderBy":  Utils.DetermineOrderKey(c.Query("orderBy"), "stars", "star_id"),
		"orderKey": Utils.DetermineOrderDirection(c.Query("orderKey")),
	}
	var stars []Models.Star
	if err := Config.DB.Debug().Raw(fmt.Sprintf(`
		SELECT DISTINCT star_id, image, name, as_character FROM stars ORDER BY %s %s
	`, params["orderBy"], params["orderKey"])).Find(&stars).Error; err == nil {
		c.JSON(http.StatusOK, stars)
	} else {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
	}
}
