package principles

import "sort"

type Result struct {
	Ticker           string
	ActualPrice      float64
	GrahanFairPrice  float64
	GrahanSafeMargin float64
	BarsiFairPrice   float64
	BarsiSafeMargin  float64
}

func GetStocks(tickers []string, dividendYield, dividendYears float64) ([]Result, error) {
	var results []Result

	for _, ticker := range tickers {
		grahan := GetGrahan(ticker)
		barsi := GetBarsi(ticker, grahan.ActualPrice, dividendYield/100, dividendYears)

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

	return results, nil
}
