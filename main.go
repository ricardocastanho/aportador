package main

import (
	"aportador/principles"
)

func main() {
	ticker := "BBAS3"

	price := principles.Grahan(ticker)
	principles.Barsi(ticker, price)
}
