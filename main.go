package main

import (
	"aportador/principles"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

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
					&cli.StringFlag{
						Name:    "tickers",
						Value:   "BBAS3,TAEE11",
						Usage:   "Stock's ticker in Brazilian stock exchange splited by comma with no spaces",
						Aliases: []string{"t"},
					},
				},
				Action: func(ctx *cli.Context) error {
					tickers := strings.Split(ctx.String("tickers"), ",")

					for _, ticker := range tickers {
						grahan := principles.GetGrahan(ticker)
						barsi := principles.GetBarsi(ticker, grahan.ActualPrice)

						fmt.Printf("Ticker: %s\n", ticker)
						fmt.Printf("Actual price: %f\n", grahan.ActualPrice)
						fmt.Printf("Grahan fair price: %f\n", grahan.FairPrice)
						fmt.Printf("Barsi fair price: %f\n", barsi.FairPrice)
						fmt.Printf("Grahan safe margin: %f\n", grahan.FairPrice)
						fmt.Printf("Barsi safe margin: %f\n", barsi.SafeMargin)
						fmt.Println("-----------------")
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
