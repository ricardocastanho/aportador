package main

import (
	"aportador/principles"
	"fmt"
)

func main() {
	ticker := "BBAS3"

	grahan := principles.GetGrahan(ticker)
	barsi := principles.GetBarsi(ticker, grahan.ActualPrice)

	fmt.Printf("Ticker: %s\n", ticker)
	fmt.Printf("Actual price: %f\n", grahan.ActualPrice)
	fmt.Printf("Grahan fair price: %f\n", grahan.FairPrice)
	fmt.Printf("Barsi fair price: %f\n", barsi.FairPrice)
	fmt.Printf("Grahan safe margin: %f\n", grahan.FairPrice)
	fmt.Printf("Barsi safe margin: %f\n", barsi.SafeMargin)
}
