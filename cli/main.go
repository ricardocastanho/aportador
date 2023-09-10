package cli

import (
	"aportador/principles"
	"aportador/server"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func CreateCLI() error {
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
						Name:    "dividend-history",
						Value:   5,
						Usage:   "Number of years to get the dividend average per year in Bazin's formula",
						Aliases: []string{"dh"},
					},
				},
				Action: func(ctx *cli.Context) error {
					tickers := ctx.StringSlice("tickers")
					dividendYield := ctx.Float64("dividend-yield")
					dividendHistory := ctx.Float64("dividend-history")

					results, err := principles.GetStocks(tickers, dividendYield, dividendHistory)

					if err != nil {
						fmt.Println("Error getting stocks data.")
						return err
					}

					PrintResults(results)
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

	return app.Run(os.Args)
}

func PrintResults(results []principles.Result) {
	for _, result := range results {
		PrintResult(result)
	}
}

func PrintResult(result principles.Result) {
	fmt.Println("-----------------")
	fmt.Printf("Ticker: %s\n", result.Ticker)
	fmt.Printf("Actual price: %f\n", result.ActualPrice)
	fmt.Printf("Grahan fair price: %f\n", result.GrahanFairPrice)
	fmt.Printf("Bazin fair price: %f\n", result.BazinFairPrice)
	fmt.Printf("Grahan safe margin: %f\n", result.GrahanSafeMargin)
	fmt.Printf("Bazin safe margin: %f\n", result.BazinSafeMargin)
	fmt.Println("-----------------")
}
