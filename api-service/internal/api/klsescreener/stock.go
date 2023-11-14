package klsescreener 

import (
	"cmp"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	BURSA_TICKER_HISTORICAL_URL = "https://www.klsescreener.com/v2/stocks/chart/"
)

type OHLC struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}

func GetBursaTickerHistoricalData(ticker string) ([]*OHLC, error) {
	var (
		ohlcs []*OHLC
	)

	url := BURSA_TICKER_HISTORICAL_URL + ticker

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("failed to create new request", "url", url, "error", err)
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("failed to get response from url", "url", url, "error", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read response body", "url", url, "error", err)
		return nil, err
	}

	// regex to remove all newlines, spaces, tabs
	reNewline := regexp.MustCompile(`\s+`)
	body = reNewline.ReplaceAll(body, []byte(""))

	reData := regexp.MustCompile(`data=\[(.*?),\];`)
	data := reData.FindAllStringSubmatch(string(body), 1)
	reData = regexp.MustCompile(`\[(.*?)\]`)
	data = reData.FindAllStringSubmatch(data[0][1], -1)

	for _, d := range data {
		split := strings.Split(d[1], ",")

		ohlc := new(OHLC)
		ohlc.Date = split[0]
		ohlc.Open = StringToFloat(split[1])
		ohlc.High = StringToFloat(split[2])
		ohlc.Low = StringToFloat(split[3])
		ohlc.Close = StringToFloat(split[4])
		ohlc.Volume = StringToFloat(split[5])
		ohlcs = append(ohlcs, ohlc)
	}

	slices.SortFunc(ohlcs, func(a, b *OHLC) int {
		dateA, _ := strconv.Atoi(a.Date)
		dateB, _ := strconv.Atoi(b.Date)
		return cmp.Compare(dateA, dateB)
	})

	return ohlcs, nil
}

func StringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		slog.Error("failed to convert string to float", "string", s, "error", err)
		return 0
	}
	return f
}
