package main

import (
	"qvarate_api/cron"
	"qvarate_api/database"
	"qvarate_api/internal"
)

func main() {

	database.GetDB()
	database.Automigrate()

	cron.SetupCron()

	s := internal.SetupRoutes()
	s.Run()
}
