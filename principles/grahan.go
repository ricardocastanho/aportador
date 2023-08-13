package principles

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	colly "github.com/gocolly/colly/v2"
)

const GRAHAN = 22.5

func Grahan(ticker string) float64 {
	var (
		raw = make(map[string][]string)
		res = make(map[string]float64)
	)

	c := colly.NewCollector()

	c.OnHTML("table tbody td.label span.txt", func(e *colly.HTMLElement) {
		raw["label"] = append(raw["label"], e.Text)
	})

	c.OnHTML("table tbody td.data span", func(e *colly.HTMLElement) {
		raw["data"] = append(raw["data"], e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		for i := range raw["data"] {
			if raw["label"][i] == "LPA" || raw["label"][i] == "VPA" || raw["label"][i] == "Cotação" {
				s := strings.ReplaceAll(raw["data"][i], ",", ".")
				n, err := strconv.ParseFloat(s, 64)

				if err != nil {
					fmt.Printf("Error during %s conversion\n", s)
					return
				}

				res[raw["label"][i]] = n
			}
		}

		grahan := math.Sqrt(GRAHAN * res["LPA"] * res["VPA"])
		safeMargin := ((grahan - res["Cotação"]) / res["Cotação"]) * 100

		fmt.Printf("Ticker: %s\n", ticker)
		fmt.Printf("Actual price: %f\n", res["Cotação"])
		fmt.Printf("Grahan fair price: %f\n", grahan)
		fmt.Printf("Grahan safe margin: %f\n", safeMargin)
	})

	c.Visit(fmt.Sprintf("https://www.fundamentus.com.br/detalhes.php?papel=%s", ticker))

	return res["Cotação"]
}
