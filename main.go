package main

import (
	"aportador/principles"
	"aportador/server"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Aportador",
		Usage: "Find Brazilian stocks best prices",
		Commands: []*cli.Command{
			{
				Name:    "stocks",
				Aliases: []string{"s"},
				Usage:   "Shows fair prices and safe margin from a stock",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "tickers",
						Value:   cli.NewStringSlice("BBAS3", "TAEE11", "PETR4"),
						Usage:   "Stocks' ticker in Brazilian stock exchange splited by comma with no spaces",
						Aliases: []string{"t"},
					},
					&cli.Float64Flag{
						Name:    "dividend-yield",
						Value:   6,
						Usage:   "Dividend yield target in Bazin's principle",
						Aliases: []string{"dy"},
					},
					&cli.Float64Flag{
						Name:    "dividend-years",
						Value:   5,
						Usage:   "Number of years to get the dividend average per year in Bazin's formula",
						Aliases: []string{"dh"},
					},
				},
				Action: func(ctx *cli.Context) error {
					tickers := ctx.StringSlice("tickers")
					dividendYield := ctx.Float64("dividend-yield")
					dividendYears := ctx.Float64("dividend-years")

					results, err := principles.GetStocks(tickers, dividendYield, dividendYears)

					if err != nil {
						fmt.Println("Error getting stocks data.")
						return err
					}

					printResults(results)
					return nil
				},
			},
			{
				Name:  "serve",
				Usage: "Listen HTTP server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "port",
						Value:   "4000",
						Usage:   "Server port",
						Aliases: []string{"p"},
					},
				},
				Action: func(ctx *cli.Context) error {
					port := ctx.String("port")
					return server.CreateServer(":" + port)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func printResults(results []principles.Result) {
	for _, result := range results {
		printResult(result)
	}
}

func printResult(result principles.Result) {
	fmt.Println("-----------------")
	fmt.Printf("Ticker: %s\n", result.Ticker)
	fmt.Printf("Actual price: %f\n", result.ActualPrice)
	fmt.Printf("Grahan fair price: %f\n", result.GrahanFairPrice)
	fmt.Printf("Bazin fair price: %f\n", result.BazinFairPrice)
	fmt.Printf("Grahan safe margin: %f\n", result.GrahanSafeMargin)
	fmt.Printf("Bazin safe margin: %f\n", result.BazinSafeMargin)
	fmt.Println("-----------------")
}
