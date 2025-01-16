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
	"time"
)

type Data struct {
	ID     string  `json:"_id"`
	Median float64 `json:"median"`
}

func GetData() {

	currency := []string{"USD", "ECU", "MLC", "CAD"}
	var newData models.Currency

	for idx, c := range currency {

		resp, err := http.Get(fmt.Sprintf(os.Getenv("DATA_API"), c)) // variable de entorno
		if err != nil {
			log.Fatal("error getting exchange rates", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("error reading data", err)
		}

		var data []Data
		if err := json.Unmarshal(body, &data); err != nil {
			log.Fatal("error unmarshaling data", err)
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
			case "CAD":
				newData.Cad = data[0].Median
			}
		}
	}

	if err := repositories.NewDateRepository().NewCurrency(newData); err != nil {
		log.Println("error saving data", err)	
	}
}
