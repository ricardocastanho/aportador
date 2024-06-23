package principles

import (
	"fmt"

	colly "github.com/gocolly/colly/v2"
)

type Bazin struct {
	Ticker string
	Price  string
	Shares string
	Profit string
	Payout string
}

func GetBazin(ticker string) Bazin {
	bazin := Bazin{
		Ticker: ticker,
	}

	var (
		raw     = make(map[string][]string)
		scraped = make(map[string]string)
	)

	c := colly.NewCollector()

	c.OnHTML("table tbody td.label span.txt", func(e *colly.HTMLElement) {
		raw["label"] = append(raw["label"], e.Text)
	})

	c.OnHTML("table tbody td.label span.txt", func(e *colly.HTMLElement) {
		raw["label"] = append(raw["label"], e.Text)
	})

	c.OnHTML("table tbody td.data span", func(e *colly.HTMLElement) {
		raw["data"] = append(raw["data"], e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		profitCount := 0
		for i := range raw["data"] {
			if raw["label"][i] == "Cotação" || raw["label"][i] == "Nro. Ações" {
				scraped[raw["label"][i]] = raw["data"][i]
			}

			if raw["label"][i] == "Lucro Líquido" && profitCount == 0 {
				scraped[raw["label"][i]] = raw["data"][i]
				profitCount++
			}
		}

		bazin.Price = scraped["Cotação"]
		bazin.Shares = scraped["Nro. Ações"]
		bazin.Profit = scraped["Lucro Líquido"]
	})

	c.Visit(fmt.Sprintf("https://www.fundamentus.com.br/detalhes.php?papel=%s", ticker))

	return bazin
}
