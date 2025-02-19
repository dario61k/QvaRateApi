package models

import "time"

type Currency struct {
	Date time.Time `gorm:"type:date;primary_key"`
	Usd  float64   `gorm:"type:decimal(10,2);not null"`
	Eur  float64   `gorm:"type:decimal(10,2);not null"`
	Mlc  float64   `gorm:"type:decimal(10,2);not null"`
}
