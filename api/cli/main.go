package main

import (
	"aportador/principles"
	"aportador/server"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	CreateCLI()
}

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
				},
				Action: func(ctx *cli.Context) error {
					tickers := ctx.StringSlice("tickers")

					results, err := principles.GetStocks(tickers)

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
	fmt.Printf("Price: %s\n", result.Price)
	fmt.Printf("Shares: %s\n", result.Shares)
	fmt.Printf("Profit: %s\n", result.Profit)
	fmt.Printf("Payout: %s\n", result.Payout)
	fmt.Println("-----------------")
}
