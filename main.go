package main

import (
	"aportador/principles"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

type Result struct {
	Ticker           string
	ActualPrice      float64
	GrahanFairPrice  float64
	GrahanSafeMargin float64
	BarsiFairPrice   float64
	BarsiSafeMargin  float64
}

func main() {
	app := &cli.App{
		Name:  "Aportador",
		Usage: "Find Brazilian stocks best prices",
		Commands: []*cli.Command{
			{
				Name:    "search",
				Aliases: []string{"s"},
				Usage:   "Shows fair prices and safe margin from a stock",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "tickers",
						Value:   cli.NewStringSlice("BBAS3", "TAEE11", "PETR4"),
						Usage:   "Stocks' ticker in Brazilian stock exchange splited by comma with no spaces",
						Aliases: []string{"t"},
					},
				},
				Action: func(ctx *cli.Context) error {
					var results []Result

					tickers := ctx.StringSlice("tickers")

					for _, ticker := range tickers {
						grahan := principles.GetGrahan(ticker)
						barsi := principles.GetBarsi(ticker, grahan.ActualPrice)

						results = append(results, Result{
							Ticker:           ticker,
							ActualPrice:      grahan.ActualPrice,
							GrahanFairPrice:  grahan.FairPrice,
							GrahanSafeMargin: grahan.SafeMargin,
							BarsiFairPrice:   barsi.FairPrice,
							BarsiSafeMargin:  barsi.SafeMargin,
						})
					}

					sort.Slice(results, func(i, j int) bool {
						if results[i].BarsiSafeMargin == results[j].BarsiSafeMargin {
							return results[i].GrahanSafeMargin > results[j].GrahanSafeMargin
						}
						return results[i].BarsiSafeMargin > results[j].BarsiSafeMargin
					})

					printResults(results)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func printResults(results []Result) {
	for _, result := range results {
		printResult(result)
	}
}

func printResult(result Result) {
	fmt.Println("-----------------")
	fmt.Printf("Ticker: %s\n", result.Ticker)
	fmt.Printf("Actual price: %f\n", result.ActualPrice)
	fmt.Printf("Grahan fair price: %f\n", result.GrahanFairPrice)
	fmt.Printf("Barsi fair price: %f\n", result.BarsiFairPrice)
	fmt.Printf("Grahan safe margin: %f\n", result.GrahanSafeMargin)
	fmt.Printf("Barsi safe margin: %f\n", result.BarsiSafeMargin)
	fmt.Println("-----------------")
}
