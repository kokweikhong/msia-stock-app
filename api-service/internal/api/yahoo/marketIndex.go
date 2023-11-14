package yahoo

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type YahooFinanceService interface {
	GetKLSEIndex() ([]*KLSEIndexHistoricalData, error)
    GetTickerHistoricalPrice(symbol string) ([]*YahooTickerHistoricalData, error)
}

type yahooFinanceService struct{}

func NewYahooFinanceService() YahooFinanceService {
	return &yahooFinanceService{}
}

const (
	KLSE_INDEX_URL = "https://query1.finance.yahoo.com/v8/finance/chart/%5EKLSE?region=MY&lang=en-US&includePrePost=false&interval=1d&useYfid=true&range=5y&corsDomain=finance.yahoo.com&.tsrc=finance"
)

type KLSEIndexResponse struct {
	Chart struct {
		Result []struct {
			Meta       interface{} `json:"meta"`
			Timestamp  []int64     `json:"timestamp"`
			Indicators struct {
				Adjclose []struct {
					Adjclose []float64 `json:"adjclose"`
				} `json:"adjclose"`
				Quote []struct {
					Open   []float64 `json:"open"`
					Close  []float64 `json:"close"`
					Volume []int64   `json:"volume"`
					High   []float64 `json:"high"`
					Low    []float64 `json:"low"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"chart"`
}

type KLSEIndexHistoricalData struct {
	Date     time.Time `json:"date"`
	Open     float64   `json:"open"`
	High     float64   `json:"high"`
	Low      float64   `json:"low"`
	Close    float64   `json:"close"`
	AdjClose float64   `json:"adjClose"`
	Volume   int64     `json:"volume"`
}

func (s *yahooFinanceService) GetKLSEIndex() ([]*KLSEIndexHistoricalData, error) {
	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, s.getKLSEIndexURLByPeriod("5y"), nil)
	if err != nil {
		slog.Error("Error creating request", "error", err)
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("origin", "https://finance.yahoo.com")
	req.Header.Set("referer", "https://finance.yahoo.com")

	res, err := client.Do(req)
	if err != nil {
		slog.Error("Error sending request", "error", err)
		return nil, err
	}

	defer res.Body.Close()

	klseIndexResponse := new(KLSEIndexResponse)
	err = decodeJSON(res.Body, klseIndexResponse)
	if err != nil {
		slog.Error("Error decoding response", "error", err)
		return nil, err
	}

	var klseIndexHistoricalData []*KLSEIndexHistoricalData
	for i, timestamp := range klseIndexResponse.Chart.Result[0].Timestamp {
		klseIndexHistoricalData = append(klseIndexHistoricalData, &KLSEIndexHistoricalData{
			Date:     time.Unix(timestamp, 0),
			Open:     klseIndexResponse.Chart.Result[0].Indicators.Quote[0].Open[i],
			High:     klseIndexResponse.Chart.Result[0].Indicators.Quote[0].High[i],
			Low:      klseIndexResponse.Chart.Result[0].Indicators.Quote[0].Low[i],
			Close:    klseIndexResponse.Chart.Result[0].Indicators.Quote[0].Close[i],
			AdjClose: klseIndexResponse.Chart.Result[0].Indicators.Adjclose[0].Adjclose[i],
			Volume:   klseIndexResponse.Chart.Result[0].Indicators.Quote[0].Volume[i],
		})
	}

	return klseIndexHistoricalData, nil

}

func (s *yahooFinanceService) getKLSEIndexURLByPeriod(period string) string {
	var period1, period2 int64

	// get 5 years before and month and day = 1 in unix
    period1 = time.Now().AddDate(-5, 0, 0).Unix()
    period2 = time.Now().AddDate(0, 0, 1).Unix()


	slog.Info("period1", "period1", period1)
	slog.Info("period2", "period2", period2)

	url := "https://query1.finance.yahoo.com/v8/finance/chart/%5EKLSE?symbol=%5EKLSE&period1=" +
		fmt.Sprint(period1) + "&period2=" + fmt.Sprint(period2) + "&useYfid=true&interval=1d&includePrePost=true&events=div%7Csplit%7Cearn&lang=en-US&region=US"
    slog.Info("url", "url", url)

    return url
}

func decodeJSON(body io.ReadCloser, v interface{}) error {
	return json.NewDecoder(body).Decode(v)
}
