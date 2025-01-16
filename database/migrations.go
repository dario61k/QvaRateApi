package database

import "qvarate_api/internal/models"

func Automigrate() {
	db.AutoMigrate(models.Currency{})
}
