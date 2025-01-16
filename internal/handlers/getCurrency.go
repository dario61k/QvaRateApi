package handlers

import (
	"net/http"
	"qvarate_api/internal/handlers/utils"
	"qvarate_api/internal/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCurrency(c *gin.Context) {

	start, err := utils.ParseDate(c.Param("startdate"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error parsing start date",
		})

		return
	}

	end, err := utils.ParseDate(c.Param("enddate"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error parsing end date",
		})

		return
	}
	var filters = map[string]time.Time{
		"startdate": start,
		"enddate":   end,
	}

	result, err := repositories.NewDateRepository().GetExchange(filters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error getting exchange rate",
		})

		return
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}
