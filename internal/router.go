package internal

import (
	h "qvarate_api/internal/handlers"
	"qvarate_api/internal/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.SetCors())

	api := r.Group("/api")

	api.GET("/get-currency/:startdate/:enddate", middleware.RateLimit(6, time.Minute*1), h.GetCurrency)
	api.GET("/get-excel/:startdate/:enddate", middleware.RateLimit(1, time.Minute*1), h.GetExcel)

	return r
}
