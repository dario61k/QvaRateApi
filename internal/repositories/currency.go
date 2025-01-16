package repositories

import (
	"errors"
	"qvarate_api/database"
	m "qvarate_api/internal/models"
	"sync"
	"time"

	"gorm.io/gorm"
)

var once sync.Once

type DateRepository struct {
	db *gorm.DB
}

var dateRepository *DateRepository

func NewDateRepository() *DateRepository {
	once.Do(func() {
		dateRepository = &DateRepository{db: database.GetDB()}
	})

	return dateRepository
}

func (r *DateRepository) NewCurrency(currency m.Currency) error {
	if result := r.db.Create(&currency); result.Error != nil {
		return result.Error
	}

	return nil
}

type Exchange struct {
	Date time.Time `json:"date"`
	Usd  float64   `json:"usd"`
	Eur  float64   `json:"eur"`
	Mlc  float64   `json:"mlc"`
	Cad  float64   `json:"cad"`
}

func (r *DateRepository) GetExchange(date map[string]time.Time) ([]Exchange, error) {

	start_value, start_ok := date["startdate"]
	end_value, end_ok := date["enddate"]

	if !start_ok || !end_ok {
		return []Exchange{}, errors.New("no date range provided")
	}

	selectField := []string{
		"date",
		"usd",
		"eur",
		"mlc",
		"cad",
	}

	query := r.db.Model(&m.Currency{})
	query = query.Select(selectField).Where("date between ? and ?", start_value, end_value)

	var exchange []Exchange
	if result := query.Find(&exchange); result.Error != nil {
		return []Exchange{}, result.Error
	}

	return exchange, nil
}
