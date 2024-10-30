package Controllers

import (
	"assignment/imdb_sql/src/Config"
	"assignment/imdb_sql/src/Models"
	"assignment/imdb_sql/src/Utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCompanies(c *gin.Context) {
	params := map[string]string{
		"orderBy":  Utils.DetermineOrderKey(c.Query("orderBy"), "companies", "id"),
		"orderKey": Utils.DetermineOrderDirection(c.Query("orderKey")),
	}
	var companies []Models.Company
	if err := Config.DB.Debug().Raw(fmt.Sprintf(`
		SELECT DISTINCT company_id, name FROM companies ORDER BY %s %s
	`, params["orderBy"], params["orderKey"])).Find(&companies).Error; err == nil {
		c.JSON(http.StatusOK, companies)
	} else {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
	}
}
