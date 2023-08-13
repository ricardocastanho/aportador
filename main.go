package main

import (
	"aportador/principles"
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
				Name:    "aportar",
				Aliases: []string{"a"},
				Usage:   "Shows fair prices and safe margin from a stock by ticket",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "ticker",
						Value:   "BBAS3",
						Usage:   "Stock's ticker in Brazilian stock exchange",
						Aliases: []string{"t"},
					},
				},
				Action: func(ctx *cli.Context) error {
					ticker := ctx.String("ticker")

					grahan := principles.GetGrahan(ticker)
					barsi := principles.GetBarsi(ticker, grahan.ActualPrice)

					fmt.Printf("Ticker: %s\n", ticker)
					fmt.Printf("Actual price: %f\n", grahan.ActualPrice)
					fmt.Printf("Grahan fair price: %f\n", grahan.FairPrice)
					fmt.Printf("Barsi fair price: %f\n", barsi.FairPrice)
					fmt.Printf("Grahan safe margin: %f\n", grahan.FairPrice)
					fmt.Printf("Barsi safe margin: %f\n", barsi.SafeMargin)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
