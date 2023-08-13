package principles

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	colly "github.com/gocolly/colly/v2"
)

func Barsi(ticker string, price float64) {
	var (
		raw     = make(map[string][]string)
		amounts = make(map[string]float64)
	)

	c := colly.NewCollector()

	c.OnHTML("table#resultado tbody tr td", func(e *colly.HTMLElement) {
		raw["data"] = append(raw["data"], e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		cont := 0
		for i := range raw["data"] {
			if cont == 4 {
				cont = 0
				continue
			}

			if cont == 3 && raw["data"][i] != "- " {
				year := fmt.Sprintf("%c%c%c%c", raw["data"][i][6], raw["data"][i][7], raw["data"][i][8], raw["data"][i][9])

				s := strings.ReplaceAll(raw["data"][i-2], ",", ".")
				amount, err := strconv.ParseFloat(s, 64)

				if err != nil {
					fmt.Printf("Error during %s conversion\n", s)
					return
				}

				amounts[year] = amounts[year] + amount
			}

			cont++
		}

		currentYear := time.Now().Year()

		var total float64

		for i := 1; i <= 5; i++ {
			total = total + amounts[strconv.Itoa(currentYear-i)]
		}

		avg := total / 5
		barsi := avg / 0.06
		safeMargin := ((barsi - price) / price) * 100

		fmt.Printf("Barsi fair price: %f\n", barsi)
		fmt.Printf("Barsi safe margin: %f\n", safeMargin)
	})

	c.Visit(fmt.Sprintf("https://www.fundamentus.com.br/proventos.php?papel=%s&tipo=2&chbAgruparAno=", ticker))
}
