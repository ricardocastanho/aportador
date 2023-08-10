package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	colly "github.com/gocolly/colly/v2"
)

var (
	label  []string
	data   []string
	values = make(map[string]float64)
)

func main() {
	ticker := "BBAS3"

	c := colly.NewCollector()

	c.OnHTML("table tbody td.label span.txt", func(e *colly.HTMLElement) {
		label = append(label, e.Text)
	})

	c.OnHTML("table tbody td.data span", func(e *colly.HTMLElement) {
		data = append(data, e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		for i := range data {
			if label[i] == "LPA" || label[i] == "VPA" {
				s := strings.ReplaceAll(data[i], ",", ".")
				n, err := strconv.ParseFloat(s, 64)

				if err != nil {
					fmt.Printf("Error during %s conversion\n", s)
					return
				}

				values[label[i]] = n
			}
		}

		grahan := math.Sqrt(22.5 * values["LPA"] * values["VPA"])

		fmt.Println("Método Grahan:")
		fmt.Printf("%s: %f\n", ticker, grahan)
	})

	c.Visit(fmt.Sprintf("https://www.fundamentus.com.br/detalhes.php?papel=%s", ticker))
}
