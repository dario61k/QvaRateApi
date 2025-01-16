package utils

import (
	"errors"
	"fmt"
	"time"
)

func ParseDate(date string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println(err)
		return parsedDate, errors.New("date parsing failed")
	}

	return parsedDate, nil
}
