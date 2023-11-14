package yahoo

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type yahooTickerResponse struct {
	Chart struct {
		Result []struct {
			Meta       interface{} `json:"meta"`
            Events     interface{} `json:"events"`
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

type YahooTickerHistoricalData struct {
	Date     time.Time `json:"date"`
	Open     float64   `json:"open"`
	High     float64   `json:"high"`
	Low      float64   `json:"low"`
	Close    float64   `json:"close"`
	AdjClose float64   `json:"adjClose"`
	Volume   int64     `json:"volume"`
}

func (y *yahooFinanceService) GetTickerHistoricalPrice(symbol string) ([]*YahooTickerHistoricalData, error) {

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, y.getTickerURL(symbol), nil)
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

	yahooTickerResponse := new(yahooTickerResponse)
	err = decodeJSON(res.Body, yahooTickerResponse)
	if err != nil {
		slog.Error("Error decoding JSON", "error", err)
		return nil, err
	}

	var yahooTickerHistoricalData []*YahooTickerHistoricalData
	for i := 0; i < len(yahooTickerResponse.Chart.Result[0].Timestamp); i++ {
		yahooTickerHistoricalData = append(yahooTickerHistoricalData, &YahooTickerHistoricalData{
			Date:     time.Unix(yahooTickerResponse.Chart.Result[0].Timestamp[i], 0),
			Open:     yahooTickerResponse.Chart.Result[0].Indicators.Quote[0].Open[i],
			High:     yahooTickerResponse.Chart.Result[0].Indicators.Quote[0].High[i],
			Low:      yahooTickerResponse.Chart.Result[0].Indicators.Quote[0].Low[i],
			Close:    yahooTickerResponse.Chart.Result[0].Indicators.Quote[0].Close[i],
			AdjClose: yahooTickerResponse.Chart.Result[0].Indicators.Adjclose[0].Adjclose[i],
			Volume:   yahooTickerResponse.Chart.Result[0].Indicators.Quote[0].Volume[i],
		})
	}

	return yahooTickerHistoricalData, nil

}

func (*yahooFinanceService) getTickerURL(symbol string) string {
	var period1, period2 int64

	period1 = time.Now().AddDate(-5, 0, 0).Unix()
	period2 = time.Now().Unix()

	url := "https://query1.finance.yahoo.com/v8/finance/chart/" + symbol + "?symbol=" + symbol + "&period1=" +
		fmt.Sprint(period1) + "&period2=" + fmt.Sprint(period2) + "&useYfid=true&interval=1d&includePrePost=true&events=div%7Csplit%7Cearn&lang=en-US&region=US"
    slog.Info("URL", "url", url)
	return url
}
