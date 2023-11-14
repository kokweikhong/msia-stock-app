package klsescreener 

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BURSA_MARKET_HISTORICAL_URL = "https://www.klsescreener.com/v2/markets/historical_period/KLSE/"
)

type BursaHistoricalData struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}

// create a enum type for year 1y or 10y
type BursaMarketYear int

const (
	OneYear BursaMarketYear = iota
	TenYear
)

func GetBursaHistoricalData(year BursaMarketYear) ([]BursaHistoricalData, error) {

	var url string
	switch year {
	case OneYear:
		url = BURSA_MARKET_HISTORICAL_URL + "1y"
	case TenYear:
		url = BURSA_MARKET_HISTORICAL_URL + "10y"
	default:
		return nil, fmt.Errorf("invalid year")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code was %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resData := struct {
		Data []BursaHistoricalData `json:"data"`
	}{}

	if err := json.Unmarshal(body, &resData); err != nil {
		return nil, err
	}

	data := resData.Data

	return data, nil

}
