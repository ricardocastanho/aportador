package principles

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	colly "github.com/gocolly/colly/v2"
)

type Bazin struct {
	Ticker        string
	DividendYield float64
	DividendYears float64
	ActualPrice   float64
	FairPrice     float64
	SafeMargin    float64
}

func GetBazin(ticker string, actualPrice, dividendYield, dividendYears float64) Bazin {
	bazin := Bazin{
		Ticker:        ticker,
		DividendYield: dividendYield,
		DividendYears: dividendYears,
		ActualPrice:   actualPrice,
	}

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
				layout := "02/01/2006"
				date, err := time.Parse(layout, strings.TrimSpace(raw["data"][i]))

				if err != nil {
					fmt.Println("Error while parsing date: ", err)
					return
				}

				year := strconv.Itoa(date.Year())

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

		for i := 1; i <= int(bazin.DividendYears); i++ {
			total = total + amounts[strconv.Itoa(currentYear-i)]
		}

		bazin.FairPrice = (total / bazin.DividendYears) / bazin.DividendYield
		bazin.SafeMargin = ((bazin.FairPrice - bazin.ActualPrice) / bazin.ActualPrice) * 100
	})

	c.Visit(fmt.Sprintf("https://www.fundamentus.com.br/proventos.php?papel=%s&tipo=2", ticker))

	return bazin
}
