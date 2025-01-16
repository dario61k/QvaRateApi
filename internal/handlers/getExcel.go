package handlers

import (
	"net/http"
	"qvarate_api/internal/handlers/utils"
	"qvarate_api/internal/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

func GetExcel(c *gin.Context) {

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
			"error": "error accessing data",
		})

		return
	}

	var data [][]interface{}
	data = append(data, []interface{}{"Fecha", "USD", "EUR", "MLC", "CAD"})

	for _, r := range result {
		data = append(data, []interface{}{r.Date, r.Usd, r.Eur, r.Mlc, r.Cad})
	}

	excelFile, err := utils.CreateExcel(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error exporting excel",
		})

		return
	}

	fileName := "tasa_de_cambio.xlsx"

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", excelFile)
}
