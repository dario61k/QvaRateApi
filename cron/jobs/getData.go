package jobs

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"qvarate_api/internal/models"
	"qvarate_api/internal/repositories"
	"sync"
	"time"
)

type Data struct {
	ID     string  `json:"_id"`
	Median float64 `json:"median"`
}

func GetData() {

	currency := []string{"USD", "ECU", "MLC"}
	var newData models.Currency

	for idx, c := range currency {

		resp, err := http.Get(fmt.Sprintf(os.Getenv("DATA_API"), c)) // variable de entorno
		if err != nil {
			log.Println("error getting exchange rates", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("error reading data", err)
		}

		var data []Data
		if err := json.Unmarshal(body, &data); err != nil {
			log.Println("error unmarshaling data", err)
		}

		if len(data) == 1 {

			if idx == 0 {
				newData.Date, _ = time.Parse("2006-01-02", data[0].ID)
			}

			switch currency[idx] {
			case "USD":
				newData.Usd = data[0].Median
			case "ECU":
				newData.Eur = data[0].Median
			case "MLC":
				newData.Mlc = data[0].Median
			}
		}
	}

	if err := repositories.NewDateRepository().NewCurrency(newData); err != nil {
		log.Println("error saving data", err)
	}
}

func GetDataV2() {

	currency := []string{"USD", "ECU", "MLC"}
	var curr models.Currency
	var m sync.Mutex
	var mw sync.WaitGroup

	for _, c := range currency {

		mw.Add(1)
		go func(c string) {

			defer mw.Done()
			resp, err := http.Get(fmt.Sprintf(os.Getenv("DATA_API"), c))
			if err != nil {
				log.Println("error getting exchange rates", err, time.Now())
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println("error reading data", err, time.Now())
			}

			var data []Data
			if err := json.Unmarshal(body, &data); err != nil {
				log.Println("error unmarshaling data", err, time.Now())
			}

			var sdata Data
			if len(data) > 0 {
				sdata = data[0]
			} else {
				log.Println("error getting data", time.Now())
				return
			}

			m.Lock()
			
			if c == "USD" {
				curr.Date, _ = time.Parse("2006-01-02", sdata.ID)
			}
			switch c {
			case "USD":
				curr.Usd = sdata.Median
			case "ECU":
				curr.Eur = sdata.Median
			case "MLC":
				curr.Mlc = sdata.Median
			}

			m.Unlock()

		}(c)
	}

	mw.Wait()

	if err := repositories.NewDateRepository().NewCurrency(curr); err != nil {
		log.Println("error saving data", err)
	}
}
