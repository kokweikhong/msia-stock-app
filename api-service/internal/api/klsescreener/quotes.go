package klsescreener

import (
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"log/slog"
)

const (
	urlQuote = "https://www.klsescreener.com/v2/screener/quote_results"
)

type Quote struct {
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	ShortName string  `json:"short_name"`
	Category  string  `json:"category"`
	Sector    string  `json:"sector"`
	EPS       float64 `json:"eps"`
	NTA       float64 `json:"nta"`
	PE        float64 `json:"pe"`
	DY        float64 `json:"dy"`
	ROE       float64 `json:"roe"`
}

func GetQuoteResults() ([]*Quote, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", urlQuote, nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("GetQuoteResults", "err", err)
		return nil, err
	}

	slog.Info("GetQuoteResults", "status", resp.Status)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("GetQuoteResults", "err", err)
		return nil, err
	}

	reTr := regexp.MustCompile(`(?s)<tr class="list">(.*?)<\/tr>`)
	trs := reTr.FindAllStringSubmatch(string(body), -1)

	slog.Info("GetQuoteResults", "trs", len(trs))

	quotes := make([]*Quote, len(trs))
	for trIndex, tr := range trs {
		reTd := regexp.MustCompile(`(?s)<td(.*?)>(.*?)<\/td>`)
		tds := reTd.FindAllStringSubmatch(tr[1], -1)

		slog.Info("GetQuoteResults", "tds", len(tds))
		reTitle := regexp.MustCompile(`(?s)<td title="(.*?)">`)

		quote := new(Quote)
		for index, td := range tds {
			title := reTitle.FindStringSubmatch(td[0])
			if len(title) < 2 {
				continue
			}
			switch title[1] {
			case "EPS":
				quote.EPS, _ = strconv.ParseFloat(strings.TrimSpace(td[1]), 64)

			}
			switch index {
			case 0:
				reA := regexp.MustCompile(`(?s)<a(.*?)>(.*?)<\/a>`)
				quote.Name = strings.TrimSpace(title[1])

				names := reA.FindStringSubmatch(td[0])
				quote.ShortName = strings.TrimSpace(names[2])
			case 1:
				codes := reTd.FindStringSubmatch(td[0])
				quote.Code = strings.TrimSpace(codes[2])
			case 2:
				reSmall := regexp.MustCompile(`(?s)<small[^>]*>(.*?)<\/small>`)
				smalls := reSmall.FindAllStringSubmatch(td[0], -1)
				quote.Category = strings.TrimSpace(smalls[0][1])
				quote.Sector = strings.TrimSpace(smalls[1][1])
			}
		}
		slog.Info("GetQuoteResults", "quote", quote)
		quotes[trIndex] = quote
	}

	slog.Info("GetQuoteResults", "total quotes", len(quotes), "total trs", len(trs))

	return quotes, nil
}
