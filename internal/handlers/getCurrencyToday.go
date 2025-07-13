package handlers

import (
	"net/http"
	"qvarate_api/internal/repositories"

	"github.com/gin-gonic/gin"
)

func GetCurrencyToday(c *gin.Context) {

	result, err := repositories.NewDateRepository().GetLastExchange()
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
